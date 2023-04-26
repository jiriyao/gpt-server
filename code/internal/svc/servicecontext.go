package svc

import (
	"chatgptserver/code/internal/config"
	"chatgptserver/code/internal/middleware"
	"github.com/r3labs/sse/v2"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config      config.Config
	CommonRoute rest.Middleware
	SseServer   *sse.Server
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		CommonRoute: middleware.NewCommonRouteMiddleware().Handle,
		SseServer:   sse.New(),
	}
}
