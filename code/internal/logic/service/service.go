package service

import (
	"chatgptserver/code/internal/svc"
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/otiai10/openaigo"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

func GetQuestionResponseAsync(client gpt3.Client, ctx context.Context, svcCtx *svc.ServiceContext, quesiton string) (string, error) {
	var sb strings.Builder

	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{ //命令行方式可以用
		Prompt: []string{
			quesiton,
		},
		MaxTokens:   gpt3.IntPtr(svcCtx.Config.ChatGpt.ApiTokenLen),
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		if len(resp.Choices) > 0 {
			for _, v := range resp.Choices {
				fmt.Print(v.Text)
				sb.WriteString(v.Text)
			}
		}
	})

	if err != nil {
		return "", err
	}

	return sb.String(), nil
}

func GetQuestionResponseAsyncCb(client gpt3.Client, ctx context.Context, svcCtx *svc.ServiceContext, quesiton string, paySuccess func(v string)) error {
	err := client.CompletionStreamWithEngine(ctx, gpt3.GPT3Dot5Turbo, gpt3.CompletionRequest{ //命令行方式可以用
		Prompt: []string{
			quesiton,
		},
		MaxTokens:   gpt3.IntPtr(svcCtx.Config.ChatGpt.ApiTokenLen),
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		if len(resp.Choices) > 0 {
			for _, v := range resp.Choices {
				paySuccess(v.Text)
			}
		}
	})

	if err != nil {
		return err
	}

	return nil
}

func GetQuestionResponseSync(client gpt3.Client, ctx context.Context, svcCtx *svc.ServiceContext, quesiton string) (string, error) {
	var sb strings.Builder

	rsp, err := client.CompletionWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			quesiton,
		},
		MaxTokens:   gpt3.IntPtr(svcCtx.Config.ChatGpt.ApiTokenLen),
		Temperature: gpt3.Float32Ptr(0),
	})

	if err != nil {
		return "", err
	}

	if len(rsp.Choices) > 0 {
		for _, v := range rsp.Choices {
			sb.WriteString(v.Text)
			break
		}
	}
	logx.Infof("gpt response text: %s \n %s \n", quesiton, sb)
	return sb.String(), nil
}

//gpt-3.5-turbo-0301
//gpt-3.5-turbo
func GetQuestionResponseSync35(ctx context.Context, svcCtx *svc.ServiceContext, quesiton string, paySuccess func(openaigo.ChatChoice)) error {
	client := openaigo.NewClient(svcCtx.Config.ChatGpt.ApiKey)
	request := openaigo.ChatCompletionRequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []openaigo.ChatMessage{
			{Role: "user", Content: quesiton},
		},
	}

	rsp, err := client.Chat(ctx, request)

	if err != nil {
		return err
	}

	if len(rsp.Choices) > 0 {
		for _, v := range rsp.Choices {
			paySuccess(v)
		}
	}
	return nil
}

func GetQuestionResponseAsyncStream(ctx context.Context, svcCtx *svc.ServiceContext, quesiton string, onData func(gpt3.ChatCompletionStreamResponseChoice)) error {
	client := gpt3.NewClient(svcCtx.Config.ChatGpt.ApiKey)
	err := client.ChatCompletionStream(ctx, gpt3.ChatCompletionRequest{ //命令行方式可以用
		Model: "gpt-3.5-turbo",
		Messages: []gpt3.ChatCompletionRequestMessage{
			{
				Role:    "system",
				Content: "You are a poetry writing assistant",
			},
			{
				Role:    "user",
				Content: quesiton,
			},
		},
		Stream:      true,
		MaxTokens:   3000,
		Temperature: 0,
		Stop:        []string{"."},
	}, func(resp *gpt3.ChatCompletionStreamResponse) {
		if len(resp.Choices) > 0 {
			for _, v := range resp.Choices {
				onData(v)
			}
		}
	})

	if err != nil {
		return err
	}

	return nil
}
