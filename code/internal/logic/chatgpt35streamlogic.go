package logic

import (
	"chatgptserver/code/internal/logic/service"
	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"
	"context"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/r3labs/sse/v2"
	"github.com/zeromicro/go-zero/core/logx"
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
		err = service.GetQuestionResponseAsyncStream(l.ctx, l.svcCtx, req.Question, func(v gpt3.ChatCompletionStreamResponseChoice) {
			if v.Delta.Content != "" {
				l.svcCtx.SseServer.Publish(req.StreamId, &sse.Event{
					Data: []byte(v.Delta.Content),
				})
			}
			if v.FinishReason != "" {
				l.Logger.Infof("v.FinishReason:" + v.FinishReason)
				l.svcCtx.SseServer.RemoveStream(req.StreamId)
			}
		})

		if err != nil {
			l.Logger.Error("v.err:" + err.Error())
			l.svcCtx.SseServer.RemoveStream(req.StreamId)
		}
	}()

	return
}
