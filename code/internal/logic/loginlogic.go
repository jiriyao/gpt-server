package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"

	"chatgptserver/code/internal/svc"
	"chatgptserver/code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (*types.Resp, error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return HttpBadRequest("参数错误"), nil
	}

	if l.svcCtx.Config.User.Username != req.Username || l.svcCtx.Config.User.Password != req.Password {
		return HttpUnauthorized("用户名或密码不正确"), nil
	}

	// ---start---
	now := time.Now().Unix()
	//accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, l.svcCtx.Config.User.Username)
	if err != nil {
		return nil, err
	}
	// ---end---
	return HttpSuccess(&LoginResponse{
		Username: req.Username,
		Token:    jwtToken,
	}), nil
}

// / login successful 后调用
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds int64, userId string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
