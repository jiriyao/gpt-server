package service

import (
	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"
	"chatgptserver/pkg/utils"
	"context"
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/r3labs/sse/v2"
	"github.com/sashabaranov/go-openai"
	uuid "github.com/satori/go.uuid"
	"io"
	"strings"
	"time"
)

type ChatCompletionMessageInfo struct { //for redis struct
	Message    openai.ChatCompletionMessage
	TokenNum   int
	CreateTime int64
}

type SendMessages struct {
	ChatList       []openai.ChatCompletionMessage
	ConversationId string
	TokenNum       int
	MaxTokenNum    int
	Mode           string
}

type (
	IChatService interface {
		Save(conversationId string, message openai.ChatCompletionMessage) error
		BuildMessageList(req *types.ChatProcessRequest) (*SendMessages, error)
		SendMessage(req *types.ChatProcessRequest, messages *SendMessages)
		CheckSensitiveWords(message *openai.ChatCompletionMessage)
	}

	ChatService struct {
		svcCtx     *svc.ServiceContext
		apiService *ApiService
	}
)

func CreateChatServiceInstance(svcCtx *svc.ServiceContext, apiService *ApiService) IChatService {
	return &ChatService{
		svcCtx:     svcCtx,
		apiService: apiService,
	}
}

func (c *ChatService) getRoleRedisKey(conversationId string) string {
	return conversationId + ":role:system"
}

//回话记录存在redis中
func (c *ChatService) Save(conversationId string, message openai.ChatCompletionMessage) error {
	//一个回话只能一个角色
	//一、设置聊天记录过期时间
	//二、更新聊天记录
	//三、聊天上下文大小控制
	//四、聊天tokens 大小控制,这个应该是发送时检测
	//把当前消息的tokens 提前计算后存在redis里
	var msg = []openai.ChatCompletionMessage{}
	msg = append(msg, message)

	tokenNum := c.apiService.NumTokensFromMessages(msg, "")

	info := ChatCompletionMessageInfo{
		Message:    message,
		TokenNum:   tokenNum,
		CreateTime: time.Now().UnixMilli(),
	}

	jsonInfo, _ := json2.Marshal(info)
	config := ChatConfigInstance.CurrentConfig(0) //通过当前vip等级获取相应的配置
	config = ChatConfigInstance.DefaultConfig()

	if message.Role == openai.ChatMessageRoleSystem {
		c.svcCtx.Redis.Setex(c.getRoleRedisKey(conversationId), string(jsonInfo), config.SaveContextTime*60)
		c.svcCtx.Redis.Expire(c.getRoleRedisKey(conversationId), config.SaveContextTime*60) //过期时间延续
		return nil
	} else {
		c.svcCtx.Redis.Lpush(conversationId, string(jsonInfo))
	}

	count, _ := c.svcCtx.Redis.Llen(conversationId)
	if count > config.ContextLength { //TODO: 这里判断保留上下文记录数量限制
		//删除多余的旧记录
		size := count - config.ContextLength
		for i := 0; i < size; i++ {
			c.svcCtx.Redis.Rpop(conversationId)
		}
	}

	//没有新的聊天记录就自动过期
	c.svcCtx.Redis.Expire(conversationId, config.SaveContextTime*60)                    //过期时间延续
	c.svcCtx.Redis.Expire(c.getRoleRedisKey(conversationId), config.SaveContextTime*60) //过期时间延续
	return nil
}

//组建发送会话消息列表
func (c *ChatService) BuildMessageList(req *types.ChatProcessRequest) (*SendMessages, error) {
	conversationId := req.ConversationId
	if conversationId == "" {
		conversationId = uuid.NewV4().String()
	}

	messageUser := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: req.Prompt,
	}

	messageSystem := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: utils.IfString(req.SystemMessage == "", c.apiService.DefaultSystemMessage(), req.SystemMessage),
	}

	c.Save(conversationId, messageUser)   //更新redis
	c.Save(conversationId, messageSystem) //更新redis

	var messages = &SendMessages{}
	messages.ConversationId = conversationId
	messageList, err := c.svcCtx.Redis.Lrange(conversationId, 0, -1)
	if err != nil {
		return nil, err
	}
	var list []openai.ChatCompletionMessage

	roleInfo, err := c.svcCtx.Redis.Get(c.getRoleRedisKey(conversationId))
	if err != nil {
		return nil, err
	}

	var numTokens = 0
	roleMsg := &ChatCompletionMessageInfo{}
	err = json2.Unmarshal([]byte(roleInfo), roleMsg)
	list = append(list, roleMsg.Message)

	numTokens += roleMsg.TokenNum
	for _, msg := range messageList {
		chatMsg := &ChatCompletionMessageInfo{}
		err = json2.Unmarshal([]byte(msg), chatMsg)
		if err != nil {
			continue
		}
		//list = append(list, chatMsg.Message)
		numTokens += chatMsg.TokenNum
		if !c.apiService.isValidPrompt(numTokens) { //如果超过了截断
			break
		}

		index := 1
		indexBackItems := append([]openai.ChatCompletionMessage{chatMsg.Message}, list[index:]...)
		list = append(list[:index], indexBackItems...)
	}

	messages.ChatList = list
	messages.TokenNum = numTokens
	messages.MaxTokenNum = c.apiService.MaxTokens(numTokens)
	return messages, nil
}

func (c *ChatService) SendMessage(req *types.ChatProcessRequest, messages *SendMessages) {
	go func() {
		config := openai.DefaultConfig(ApiKeyFactory.GetKey())
		client := openai.NewClientWithConfig(config)
		ctx := context.Background()
		req.ConversationId = messages.ConversationId
		reqGpt := openai.ChatCompletionRequest{
			Model:       utils.IfString(messages.Mode == "", openai.GPT3Dot5Turbo, messages.Mode),
			MaxTokens:   messages.MaxTokenNum,
			Messages:    messages.ChatList,
			Temperature: utils.IfFloat32(req.Temperature <= 0, 0.8, req.Temperature),
			TopP:        utils.IfFloat32(req.TopP <= 0, 1, req.TopP),
			Stream:      true,
		}

		if 1 == 1 {
			info, _ := json2.Marshal(reqGpt)
			fmt.Printf("ChatCompletionStream reqGpt Body: %v\n", string(info))
		}

		stream, err := client.CreateChatCompletionStream(ctx, reqGpt)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
			c.svcCtx.SseServer.RemoveStream(req.StreamId)
			return
		}
		defer stream.Close()

		fmt.Printf("Stream response: ")
		finished := false
		errored := false
		ss := strings.Builder{}
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("\nStream finished")
				finished = true
				c.sendMessageToClient(req, &response, "stop")
				c.svcCtx.SseServer.RemoveStream(req.StreamId)
				break
			}

			if err != nil {
				errored = true
				fmt.Printf("\nStream error: %v\n", err)
				c.sendMessageToClient(req, &response, "stop")
				c.svcCtx.SseServer.RemoveStream(req.StreamId)
				break
			}

			content := response.Choices[0].Delta.Content
			ss.WriteString(content)

			if content != "" {
				c.sendMessageToClient(req, &response, "")
			}
		}

		if errored {
			return
		}

		if finished {
			c.Save(req.ConversationId, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: ss.String(),
			})
		}
		return
	}()
}

func (c *ChatService) sendMessageToClient(req *types.ChatProcessRequest, response *openai.ChatCompletionStreamResponse, finishReason string) {
	resp := buildResponseText(req, response, finishReason)

	json, _ := json2.Marshal(resp)
	c.svcCtx.SseServer.Publish(req.StreamId, &sse.Event{
		Data: json,
	})

	//c.svcCtx.SseServer.Publish(req.StreamId, &sse.Event{
	//	Data: []byte("test"),
	//})
}

func buildResponseText(req *types.ChatProcessRequest, response *openai.ChatCompletionStreamResponse, finishReason string) *types.ChatProcessResponse {
	content := ""
	if finishReason != "stop" {
		content = response.Choices[0].Delta.Content
	} else {
		//TODO 整个聊天记录存在Redis
		content = ""
	}
	resp := types.ChatProcessResponse{
		Role:            openai.ChatMessageRoleAssistant,
		Model:           response.Model,
		Id:              response.ID,
		ParentMessageId: req.ConversationId,
		DeltaText:       content,
		FinishReason:    finishReason,
	}
	return &resp
}

//

// 发送问题合法性判断
func (c *ChatService) CheckRecord(message *openai.ChatCompletionMessage) bool {
	return false
}

// 发送问题合法性判断，并只保留合法长度
func (c *ChatService) CutRecord(message *openai.ChatCompletionMessage) {

}

func (c *ChatService) CheckSensitiveWords(message *openai.ChatCompletionMessage) {

}
