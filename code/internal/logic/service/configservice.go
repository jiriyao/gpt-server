package service

import "chatgptserver/pkg/utils"

type GlobalConfig struct { //管理端配置
	TotalPrompts       int //总共请求数量
	VipTotalPrompts    int
	TotalToken         int //总共tokens 数量限制
	VipToken           int
	MaxToken           int
	VipMaxToken        int
	ContextLength      int //保留上下文的最近记录数量
	VipContextLength   int //上下问
	SaveContextTime    int //保留记录时间，策略是：如果一直聊不会自动删除，如果过了一这个时间后自动清除，再聊就是新话,分钟
	VipSaveContextTime int //保留记录时间，策略是：如果一直聊不会自动删除，如果过了一这个时间后自动清除，再聊就是新话
}

type CurrentConfig struct {
	Conversations   int //同时能开几个回话
	TotalPrompts    int //总共请求问题数量（请求次数）
	TotalToken      int //总共请求token数量
	MaxToken        int //一次请求最大token
	ContextLength   int //一个会话保留上下文的最近记录数量
	SaveContextTime int //一个会话保留记录时间，策略是：如果一直聊不会自动删除，如果过了一这个时间后自动清除，再聊就是新话,分钟
}

type ConfigService struct {
}

var ChatConfigInstance = ConfigService{}

func (s *ConfigService) GlobalConfig() *GlobalConfig {
	return &GlobalConfig{
		MaxToken:           1000,
		VipMaxToken:        2000,
		ContextLength:      5,
		VipContextLength:   10,
		SaveContextTime:    30,
		VipSaveContextTime: 120,
	}
}

func (s *ConfigService) CurrentConfig(vipLevel int) *CurrentConfig {
	conf := s.GlobalConfig()
	return &CurrentConfig{
		MaxToken:        utils.IfInt(vipLevel <= VIP0, conf.MaxToken, conf.VipMaxToken),
		ContextLength:   utils.IfInt(vipLevel <= VIP0, conf.ContextLength, conf.VipContextLength),
		SaveContextTime: utils.IfInt(vipLevel <= VIP0, conf.SaveContextTime, conf.VipSaveContextTime),
	}
}

func (s *ConfigService) DefaultConfig() *CurrentConfig {
	return &CurrentConfig{
		MaxToken:        1000,
		ContextLength:   5,
		SaveContextTime: 2,
	}
}
