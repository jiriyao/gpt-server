package logic

import (
	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"
	"context"
	"errors"
	"fmt"
	"github.com/r3labs/sse/v2"
	"github.com/sashabaranov/go-openai"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
)

type Chatgpt35StreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatgpt35StreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Chatgpt35StreamLogic {
	return &Chatgpt35StreamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Chatgpt35StreamLogic) Chatgpt35Stream(req *types.ChatGptQuestionRequest) (resp *types.Resp, err error) {
	go func() {
		c := openai.NewClient(l.svcCtx.Config.ChatGpt.ApiKey)
		ctx := context.Background()

		reqGpt := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 200,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: req.Question,
				},
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You ara chatgpt.",
				},
			},
			Stream: true,
		}
		stream, err := c.CreateChatCompletionStream(ctx, reqGpt)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
			return
		}
		defer stream.Close()

		fmt.Printf("Stream response: ")
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("\nStream finished")
				l.svcCtx.SseServer.RemoveStream(req.StreamId)
				return
			}

			if err != nil {
				fmt.Printf("\nStream error: %v\n", err)
				l.svcCtx.SseServer.Publish(req.StreamId, &sse.Event{
					Data: []byte("chat error"),
				})
				l.svcCtx.SseServer.RemoveStream(req.StreamId)
				return
			}

			content := response.Choices[0].Delta.Content
			fmt.Printf(content)

			if content != "" {
				l.svcCtx.SseServer.Publish(req.StreamId, &sse.Event{
					Data: []byte(content),
				})
			}
		}
	}()

	return
}
