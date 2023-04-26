package logic

import (
	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type CodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type ChatGptRequest struct {
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float64 `json:"temperature"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
	Model            string  `json:"model"`
}

func NewCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CodeLogic {
	return &CodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//for test
func (l *CodeLogic) Code(req *types.Request) (resp *types.Resp, err error) {
	return
}
