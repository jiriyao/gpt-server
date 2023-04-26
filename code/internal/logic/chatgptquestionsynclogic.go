package logic

import (
	"chatgptserver/code/internal/logic/service"
	"context"
	"github.com/PullRequestInc/go-gpt3"

	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatgptQuestionSyncLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatgptQuestionSyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatgptQuestionSyncLogic {
	return &ChatgptQuestionSyncLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatgptQuestionSyncLogic) ChatgptQuestionSync(req *types.ChatGptQuestionRequest) (resp *types.Resp, err error) {
	client := gpt3.NewClient(l.svcCtx.Config.ChatGpt.ApiKey)
	rsp, e2 := service.GetQuestionResponseSync(client, l.ctx, l.svcCtx, req.Question)
	if e2 != nil {
		return HttpServerError(e2.Error()), nil
	}
	return HttpSuccess(types.ChatGptQuestionResponse{
		Q:      req.Question,
		Answer: rsp,
	}), nil
}
