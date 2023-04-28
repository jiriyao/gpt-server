package redis

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type Conf struct {
	cache.NodeConf
	Prefix string
}

type ClusterConf []Conf

func (c ClusterConf) GetCacheConf() cache.CacheConf {
	cacheConf := make([]cache.NodeConf, 0, len(c))
	for _, conf := range c {
		cacheConf = append(cacheConf, conf.NodeConf)
	}
	return cacheConf
}
