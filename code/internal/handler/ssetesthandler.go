package handler

import (
	"chatgptserver/code/internal/logic"
	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func SseTestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatGptQuestionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		uuid := uuid.NewV4().String()
		req.StreamId = uuid
		l := logic.NewSseTestLogic(r.Context(), svcCtx)
		_, err := l.SseTest(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		sseHandler(w, r, svcCtx, uuid)

		// 设置响应头 for testing
		//w.Header().Set("Content-Type", "text/event-stream")
		//w.Header().Set("Cache-Control", "no-cache")
		//w.Header().Set("Connection", "keep-alive")
		//
		//// 向客户端发送初始数据
		//fmt.Fprintf(w, "data: %s\n\n", "Hello SSE!")
		//
		//// 创建定时器，每 3 秒发送一条消息
		//
		//aa := make(chan string)
		//for true {
		//	println("hello aaaa")
		//	time.Sleep(2 * time.Second)
		//	data := fmt.Sprintf("data: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))
		//	fmt.Fprintf(w, data)
		//	//flusher.Flush()
		//
		//	w.(http.Flusher).Flush()
		//}
		//
		//<-aa
		//
		//println("hello qqqqq")
		////<-r.Context().Done()

	}

}
