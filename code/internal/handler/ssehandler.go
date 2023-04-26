package handler

import (
	"chatgptserver/code/internal/svc"
	"net/http"
)

func sseHandler(w http.ResponseWriter, r *http.Request, svcCtx *svc.ServiceContext, streamId string) {
	//println("The client is streamId:" + u1 + ":" + r.Method)
	go func() {
		// Received Browser Disconnection
		<-r.Context().Done()
		svcCtx.SseServer.RemoveStream(streamId)
		println("The client is disconnected here:channel:" + streamId)
		return
	}()

	//tw := w.(*http.timeoutWriter)
	//w.ResponseWriter

	//refletct.ValueOf(*w).FieldByName("Name")

	query := r.URL.Query()
	query.Set("stream", streamId)
	r.URL.RawQuery = query.Encode()
	svcCtx.SseServer.ServeHTTP(w, r) //阻塞并等待消息和断开连接
}
