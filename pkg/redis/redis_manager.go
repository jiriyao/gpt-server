package redis

import "sync"

func NewManager() *Manager {
	return &Manager{}
}

type Manager struct {
	redisClients sync.Map
}

func (m *Manager) AddRedis(redis *XzRedis) {
	m.redisClients.Store(redis.KeyPrefix, redis)
}

func (m *Manager) GetRedisBy(redis *XzRedis, prefix string) *XzRedis {
	val, ok := m.redisClients.Load(prefix)
	if ok {
		return val.(*XzRedis)
	}
	newRedis := &XzRedis{
		Redis:     redis.Redis,
		KeyPrefix: prefix,
	}

	m.redisClients.Store(prefix, newRedis)

	return newRedis
}
