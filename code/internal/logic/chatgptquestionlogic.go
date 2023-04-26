package logic

import (
	"chatgptserver/code/internal/logic/service"
	"context"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatgptQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatgptQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatgptQuestionLogic {
	return &ChatgptQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatgptQuestionLogic) ChatgptQuestion(req *types.ChatGptQuestionRequest) (resp *types.Resp, err error) {
	client := gpt3.NewClient(l.svcCtx.Config.ChatGpt.ApiKey)
	rsp, e2 := service.GetQuestionResponseAsync(client, l.ctx, l.svcCtx, req.Question)
	return HttpSuccess(rsp), e2
}

func (l *ChatgptQuestionLogic) ChatgptQuestionTest(req *types.ChatGptQuestionRequest, w http.ResponseWriter) (resp *types.Resp, err error) {
	client := gpt3.NewClient(l.svcCtx.Config.ChatGpt.ApiKey)
	e2 := service.GetQuestionResponseAsyncCb(client, l.ctx, l.svcCtx, req.Question, func(v string) {
		w.Write([]byte(v))
	})

	if e2 != nil {
		logx.Error(e2)
		httpx.ErrorCtx(l.ctx, w, err)
	}
	return nil, nil
}
