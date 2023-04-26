package main

import (
	"chatgptserver/code/internal/config"
	"chatgptserver/code/internal/handler"
	"chatgptserver/code/internal/svc"
	"flag"
	"fmt"
	"github.com/r3labs/sse/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"strings"
)

var configFile = flag.String("f", "etc/code-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	logx.DisableStat()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	logx.DisableStat()
	//go func() {
	//	sseServerStart(ctx)
	//}()
	ctx.SseServer.AutoStream = true
	server.Start()

}

func sseServerStart(ctx *svc.ServiceContext) {
	server := sse.New()
	server.AutoStream = true
	mux := http.NewServeMux()

	mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if strings.ToUpper(r.Method) != "POST" {
			w.Write([]byte("error"))
			return
		}

		u1 := uuid.NewV4().String()
		println("The client is uuid:" + u1 + ":" + r.Method)

		go func() {
			// Received Browser Disconnection
			<-r.Context().Done()
			server.RemoveStream(u1)
			println("The client is disconnected here:channel:" + u1)
			return
		}()

		query := r.URL.Query()
		query.Set("stream", u1)
		r.URL.RawQuery = query.Encode()
		server.ServeHTTP(w, r) //阻塞并等待消息和断开连接

	})

	fmt.Printf("Starting sse server at %s:%d...\n", "", "8800")
	http.ListenAndServe(":8800", mux)
}
