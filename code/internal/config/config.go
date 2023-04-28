package config

import (
	"chatgptserver/pkg/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	ChatGpt struct {
		Api              string
		ApiKey           string
		Organization     string
		ApiTokenLen      int
		AccessToken      string
		SessionToken     string
		RefreshTokenTime int64
	}

	Redis redis.ClusterConf

	User struct { //用jwt
		Username string
		Password string
	}

	Net struct {
		AesKey string
		Iv     string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
