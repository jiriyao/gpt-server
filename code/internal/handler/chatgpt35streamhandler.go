package handler

import (
	uuid "github.com/satori/go.uuid"
	"net/http"

	"chatgptserver/code/internal/logic"
	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Chatgpt35StreamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatGptQuestionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		uuid := uuid.NewV4().String()
		req.StreamId = uuid

		l := logic.NewChatgpt35StreamLogic(r.Context(), svcCtx)
		_, err := l.Chatgpt35Stream(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		sseHandler(w, r, svcCtx, uuid)
	}
}
