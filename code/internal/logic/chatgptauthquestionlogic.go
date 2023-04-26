package logic

import (
	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatgptAuthQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatgptAuthQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatgptAuthQuestionLogic {
	return &ChatgptAuthQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatgptAuthQuestionLogic) ChatgptAuthQuestion(req *types.ChatGptQuestionRequest) (resp *types.Resp, err error) {
	//chatgpt := chatgpt.NewChatGpt(chatgpt.NewClient(&chatgpt.Credentials{
	//	BearerToken:  "Bearer " + l.svcCtx.Config.ChatGpt.AccessToken,
	//	SessionToken: l.svcCtx.Config.ChatGpt.SessionToken,
	//}))
	//mockRequest := req.Question
	//
	//// Run test
	//res, err := chatgpt.SendMessage(mockRequest)
	//if err != nil {
	//	return HttpForbidden(err.Error()), nil
	//}
	//
	//if err != nil {
	//	return HttpServerError(err.Error()), nil
	//}
	//return HttpSuccess(types.ChatGptQuestionResponse{
	//	Q:      req.Question,
	//	Answer: res,
	//}), nil

	return
}
