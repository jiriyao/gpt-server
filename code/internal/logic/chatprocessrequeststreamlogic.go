package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/r3labs/sse/v2"
	"github.com/sashabaranov/go-openai"
	"io"

	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatProcessRequestStreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatProcessRequestStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatProcessRequestStreamLogic {
	return &ChatProcessRequestStreamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatProcessRequestStreamLogic) ChatProcessRequestStream(req *types.ChatProcessRequest) (resp *types.Resp, err error) {
	go func() {
		config := openai.DefaultConfig(l.svcCtx.Config.ChatGpt.ApiKey)
		//config.BaseURL = l.svcCtx.Config.ChatGpt.Api
		c := openai.NewClientWithConfig(config)
		ctx := context.Background()

		reqGpt := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 200,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: req.Prompt,
				},
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: req.SystemMessage,
				},
			},
			Temperature: req.Temperature,
			TopP:        req.TopP,
			Stream:      true,
		}
		stream, err := c.CreateChatCompletionStream(ctx, reqGpt)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
			l.svcCtx.SseServer.RemoveStream(req.StreamId)
			return
		}
		defer stream.Close()

		fmt.Printf("Stream response: ")
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("\nStream finished")
				l.sendMessageToClient(req, &response, "stop")
				l.svcCtx.SseServer.RemoveStream(req.StreamId)
				return
			}

			if err != nil {
				fmt.Printf("\nStream error: %v\n", err)
				l.sendMessageToClient(req, &response, "stop")
				l.svcCtx.SseServer.RemoveStream(req.StreamId)
				return
			}

			content := response.Choices[0].Delta.Content
			fmt.Printf(content)

			if content != "" {
				l.sendMessageToClient(req, &response, "")
			}
		}
	}()

	return
}

func (l *ChatProcessRequestStreamLogic) sendMessageToClient(req *types.ChatProcessRequest, response *openai.ChatCompletionStreamResponse, finishReason string) {
	//resp := buildResponseText(response, finishReason)
	//
	//json, _ := json2.Marshal(resp)
	//l.svcCtx.SseServer.Publish(req.StreamId, &sse.Event{
	//	Data: json,
	//})

	l.svcCtx.SseServer.Publish(req.StreamId, &sse.Event{
		Data: []byte("test"),
	})
}

func buildResponseText(response *openai.ChatCompletionStreamResponse, finishReason string) *types.ChatProcessResponse {
	content := ""
	if finishReason == "stop" {
		content = response.Choices[0].Delta.Content
	}
	resp := types.ChatProcessResponse{
		Role:            openai.ChatMessageRoleAssistant,
		Model:           response.Model,
		Id:              response.ID,
		ParentMessageId: response.ID,
		DeltaText:       content,
		FinishReason:    finishReason,
	}
	return &resp
}
