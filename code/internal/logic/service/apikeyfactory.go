package service

type ApiKeyService struct {
}

var ApiKeyFactory = ApiKeyService{}

func (s *ApiKeyService) GetKey() string {
	return "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"
}

//如果出现接口返回 401，则把当前的key 禁止使用
func (s *ApiKeyService) DisableKey(key string) {

}

func (s *ApiKeyService) EnableKey(key string) {

}

func (s *ApiKeyService) FreeKeys() []string {
	return []string{}
}

func (s *ApiKeyService) IsFree() bool {
	return false
}

//是否是一个过期的key
func (s *ApiKeyService) isDisabledKey(key string) bool {
	return false
}
