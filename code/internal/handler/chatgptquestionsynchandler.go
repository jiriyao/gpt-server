package handler

import (
	"net/http"

	"chatgptserver/code/internal/logic"
	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatgptQuestionSyncHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatGptQuestionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewChatgptQuestionSyncLogic(r.Context(), svcCtx)
		resp, err := l.ChatgptQuestionSync(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
