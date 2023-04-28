package service

const (
	VIP0 = 0
	VIP1 = 1
)

type UserService struct {
}

var UserServiceInstance = &UserService{}

//根据用户token
func (u *UserService) GetUserLevel(userToken string) int {
	if userToken == "" {
		return VIP0
	}
	//检查redis 里的用户信息
	return VIP0
}

func (u *UserService) CheckUser(userToken string) bool {
	//判断当前用户是否合法权益，tokens 请求数量，vip 过期等等 ，可以在中间件处理？？？？
	return true
}
