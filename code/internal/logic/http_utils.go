package logic

import (
	"chatgptserver/code/internal/types"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Response struct{}

// 请求成功
func HttpSuccess(data interface{}) *types.Resp {
	return &types.Resp{
		Code: 0,
		Msg:  "ok",
		Data: data,
	}
}

// 请求格式错误，比如参数格式、参数字段名等 不正确
func HttpBadRequest(msg string) *types.Resp {
	return &types.Resp{
		Code: 400,
		Msg:  msg,
		Data: Response{},
	}
}

// 用户没有访问权限，需要进行身份认证
func HttpUnauthorized(msg string) *types.Resp {
	return &types.Resp{
		Code: 401,
		Msg:  msg,
		Data: Response{},
	}
}

// 用户已进行身份认证，但权限不够
func HttpForbidden(msg string) *types.Resp {
	return &types.Resp{
		Code: 403,
		Msg:  msg,
		Data: Response{},
	}
}

// 接口不存在
func HttpNotFound(msg string) *types.Resp {
	return &types.Resp{
		Code: 404,
		Msg:  msg,
		Data: Response{},
	}
}

// 服务器内部错误
func HttpServerError(msg string) *types.Resp {
	return &types.Resp{
		Code: 500,
		Msg:  msg,
		Data: Response{},
	}
}

// 请求失败
func HttpFail(msg string) *types.Resp {
	return &types.Resp{
		Code: 10001,
		Msg:  msg,
		Data: Response{},
	}
}

// 如需返回特殊错误码，调用此接口
func HttpFailForCode(code int64, msg string) *types.Resp {
	return &types.Resp{
		Code: code,
		Msg:  msg,
		Data: Response{},
	}
}

func PostForm(urlPath string, data map[string]string) []byte {
	var param = url.Values{}
	for k, v := range data {
		param.Set(k, v)
	}
	client := &http.Client{Timeout: 20 * time.Second}
	rsp, err := client.PostForm(urlPath, param)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	r, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	return r
}

func httpPost(url string, postParams string) ([]byte, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postParams))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	return body, err
}
