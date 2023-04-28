package logic

import (
	"chatgptserver/code/internal/logic/service"
	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"

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
		chatService := service.CreateChatServiceInstance(l.svcCtx, service.WithDefaultConfig(openai.GPT3Dot5Turbo))
		msgs, er := chatService.BuildMessageList(req)
		if er != nil {
			fmt.Printf("ChatProcessRequestStream error: %v\n", err)
			l.svcCtx.SseServer.RemoveStream(req.StreamId)
		}
		chatService.SendMessage(req, msgs)
	}()
	return
}
