package handler

import (
	"net/http"

	"chatgptserver/code/internal/logic"
	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCodeLogic(r.Context(), svcCtx)
		resp, err := l.Code(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
