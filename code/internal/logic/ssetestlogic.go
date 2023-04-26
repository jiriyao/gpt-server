package logic

import (
	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"
	"context"
	"github.com/r3labs/sse/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type SseTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSseTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SseTestLogic {
	return &SseTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SseTestLogic) SseTest(req *types.ChatGptQuestionRequest) (resp *types.Resp, err error) {
	go func() {
		for true {
			time.Sleep(time.Second * 1)
			l.svcCtx.SseServer.Publish(req.StreamId, &sse.Event{
				Data: []byte("ping"),
			})
		}
	}()

	return
}
