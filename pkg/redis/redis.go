package redis

import (
	"time"

	"github.com/gogf/gf/container/garray"
	gzRedis "github.com/zeromicro/go-zero/core/stores/redis"
)

var (
	Nil = gzRedis.Nil
)

type XzRedis struct {
	//redis 实例，禁止在外部直接调用此实例的方法！
	Redis     *gzRedis.Redis
	KeyPrefix string
}

// NewRedis returns a XzRedis.
func NewRedis(conf Conf) *XzRedis {
	prefix := "unknown"
	if conf.Prefix != "" {
		prefix = conf.Prefix
	}
	return &XzRedis{
		Redis:     conf.NewRedis(),
		KeyPrefix: prefix,
	}
}

func (s *XzRedis) WrapKey(key string) string {
	return s.KeyPrefix + ":" + key
}

func (s *XzRedis) WrapKeys(keys ...string) []string {
	arr := garray.NewStrArrayFromCopy(keys)
	return arr.Walk(func(val string) string { return s.WrapKey(val) }).Slice()
}

// BitCount is redis bitcount command implementation.
func (s *XzRedis) BitCount(key string, start, end int64) (val int64, err error) {
	return s.Redis.BitCount(s.WrapKey(key), start, end)
}

// BitOpAnd is redis bit operation (and) command implementation.
func (s *XzRedis) BitOpAnd(destKey string, keys ...string) (val int64, err error) {
	return s.Redis.BitOpAnd(s.WrapKey(destKey), s.WrapKeys(keys...)...)
}

// BitOpNot is redis bit operation (not) command implementation.
func (s *XzRedis) BitOpNot(destKey, key string) (val int64, err error) {
	return s.Redis.BitOpNot(s.WrapKey(destKey), s.WrapKey(key))
}

// BitOpOr is redis bit operation (or) command implementation.
func (s *XzRedis) BitOpOr(destKey string, keys ...string) (val int64, err error) {
	return s.Redis.BitOpOr(s.WrapKey(destKey), s.WrapKeys(keys...)...)
}

// BitOpXor is redis bit operation (xor) command implementation.
func (s *XzRedis) BitOpXor(destKey string, keys ...string) (val int64, err error) {
	return s.Redis.BitOpXor(s.WrapKey(destKey), s.WrapKeys(keys...)...)
}

// BitPos is redis bitpos command implementation.
func (s *XzRedis) BitPos(key string, bit, start, end int64) (val int64, err error) {
	return s.Redis.BitPos(s.WrapKey(key), bit, start, end)
}

// Blpop uses passed in redis connection to execute blocking queries.
// Doesn't benefit from pooling redis connections of blocking queries
func (s *XzRedis) Blpop(redisNode gzRedis.RedisNode, key string) (string, error) {
	return s.Redis.Blpop(redisNode, s.WrapKey(key))
}

// BlpopEx uses passed in redis connection to execute blpop command.
// The difference against Blpop is that this method returns a bool to indicate success.
func (s *XzRedis) BlpopEx(redisNode gzRedis.RedisNode, key string) (string, bool, error) {
	return s.Redis.BlpopEx(redisNode, s.WrapKey(key))
}

// Del deletes keys.
func (s *XzRedis) Del(keys ...string) (val int, err error) {
	return s.Redis.Del(s.WrapKeys(keys...)...)
}

// Eval is the implementation of redis eval command.
func (s *XzRedis) Eval(script string, keys []string, args ...interface{}) (val interface{}, err error) {
	return s.Redis.Eval(script, s.WrapKeys(keys...), args...)
}

// EvalSha is the implementation of redis evalsha command.
func (s *XzRedis) EvalSha(sha string, keys []string, args ...interface{}) (val interface{}, err error) {
	return s.Redis.EvalSha(sha, s.WrapKeys(keys...), args...)
}

// Exists is the implementation of redis exists command.
func (s *XzRedis) Exists(key string) (val bool, err error) {
	return s.Redis.Exists(s.WrapKey(key))
}

// Expire is the implementation of redis expire command.
func (s *XzRedis) Expire(key string, seconds int) error {
	return s.Redis.Expire(s.WrapKey(key), seconds)
}

// Expireat is the implementation of redis expireat command.
func (s *XzRedis) Expireat(key string, expireTime int64) error {
	return s.Redis.Expireat(s.WrapKey(key), expireTime)
}

// GeoAdd is the implementation of redis geoadd command.
func (s *XzRedis) GeoAdd(key string, geoLocation ...*gzRedis.GeoLocation) (val int64, err error) {
	return s.Redis.GeoAdd(s.WrapKey(key), geoLocation...)
}

// GeoDist is the implementation of redis geodist command.
func (s *XzRedis) GeoDist(key, member1, member2, unit string) (val float64, err error) {
	return s.Redis.GeoDist(s.WrapKey(key), member1, member2, unit)
}

// GeoHash is the implementation of redis geohash command.
func (s *XzRedis) GeoHash(key string, members ...string) (val []string, err error) {
	return s.Redis.GeoHash(s.WrapKey(key), members...)
}

// GeoRadius is the implementation of redis georadius command.
func (s *XzRedis) GeoRadius(key string, longitude, latitude float64, query *gzRedis.GeoRadiusQuery) (val []gzRedis.GeoLocation, err error) {
	return s.Redis.GeoRadius(s.WrapKey(key), longitude, latitude, query)
}

// GeoRadiusByMember is the implementation of redis georadiusbymember command.
func (s *XzRedis) GeoRadiusByMember(key, member string, query *gzRedis.GeoRadiusQuery) (val []gzRedis.GeoLocation, err error) {
	return s.Redis.GeoRadiusByMember(s.WrapKey(key), member, query)
}

// GeoPos is the implementation of redis geopos command.
func (s *XzRedis) GeoPos(key string, members ...string) (val []*gzRedis.GeoPos, err error) {
	return s.Redis.GeoPos(s.WrapKey(key), members...)
}

// Get is the implementation of redis get command.
func (s *XzRedis) Get(key string) (val string, err error) {
	return s.Redis.Get(s.WrapKey(key))
}

// GetBit is the implementation of redis getbit command.
func (s *XzRedis) GetBit(key string, offset int64) (val int, err error) {
	return s.Redis.GetBit(s.WrapKey(key), offset)
}

// Hdel is the implementation of redis hdel command.
func (s *XzRedis) Hdel(key string, fields ...string) (val bool, err error) {
	return s.Redis.Hdel(s.WrapKey(key), fields...)
}

// Hexists is the implementation of redis hexists command.
func (s *XzRedis) Hexists(key, field string) (val bool, err error) {
	return s.Redis.Hexists(s.WrapKey(key), field)
}

// Hget is the implementation of redis hget command.
func (s *XzRedis) Hget(key, field string) (val string, err error) {
	return s.Redis.Hget(s.WrapKey(key), field)
}

// Hgetall is the implementation of redis hgetall command.
func (s *XzRedis) Hgetall(key string) (val map[string]string, err error) {
	return s.Redis.Hgetall(s.WrapKey(key))
}

// Hincrby is the implementation of redis hincrby command.
func (s *XzRedis) Hincrby(key, field string, increment int) (val int, err error) {
	return s.Redis.Hincrby(s.WrapKey(key), field, increment)
}

// Hkeys is the implementation of redis hkeys command.
func (s *XzRedis) Hkeys(key string) (val []string, err error) {
	return s.Redis.Hkeys(s.WrapKey(key))
}

// Hlen is the implementation of redis hlen command.
func (s *XzRedis) Hlen(key string) (val int, err error) {
	return s.Redis.Hlen(s.WrapKey(key))
}

// Hmget is the implementation of redis hmget command.
func (s *XzRedis) Hmget(key string, fields ...string) (val []string, err error) {
	return s.Redis.Hmget(s.WrapKey(key), fields...)
}

// Hset is the implementation of redis hset command.
func (s *XzRedis) Hset(key, field, value string) error {
	return s.Redis.Hset(s.WrapKey(key), field, value)
}

// Hsetnx is the implementation of redis hsetnx command.
func (s *XzRedis) Hsetnx(key, field, value string) (val bool, err error) {
	return s.Redis.Hsetnx(s.WrapKey(key), field, value)
}

// Hmset is the implementation of redis hmset command.
func (s *XzRedis) Hmset(key string, fieldsAndValues map[string]string) error {
	return s.Redis.Hmset(s.WrapKey(key), fieldsAndValues)
}

// Hscan is the implementation of redis hscan command.
func (s *XzRedis) Hscan(key string, cursor uint64, match string, count int64) (keys []string, cur uint64, err error) {
	return s.Redis.Hscan(s.WrapKey(key), cursor, match, count)
}

// Hvals is the implementation of redis hvals command.
func (s *XzRedis) Hvals(key string) (val []string, err error) {
	return s.Redis.Hvals(s.WrapKey(key))
}

// Incr is the implementation of redis incr command.
func (s *XzRedis) Incr(key string) (val int64, err error) {
	return s.Redis.Incr(s.WrapKey(key))
}

// Incrby is the implementation of redis incrby command.
func (s *XzRedis) Incrby(key string, increment int64) (val int64, err error) {
	return s.Redis.Incrby(s.WrapKey(key), increment)
}

// Keys is the implementation of redis keys command.
func (s *XzRedis) Keys(pattern string) (val []string, err error) {
	return s.Redis.Keys(pattern)
}

// Llen is the implementation of redis llen command.
func (s *XzRedis) Llen(key string) (val int, err error) {
	return s.Redis.Llen(s.WrapKey(key))
}

// Lpop is the implementation of redis lpop command.
func (s *XzRedis) Lpop(key string) (val string, err error) {
	return s.Redis.Lpop(s.WrapKey(key))
}

// Lpush is the implementation of redis lpush command.
func (s *XzRedis) Lpush(key string, values ...interface{}) (val int, err error) {
	return s.Redis.Lpush(s.WrapKey(key), values...)
}

// Lrange is the implementation of redis lrange command.
func (s *XzRedis) Lrange(key string, start, stop int) (val []string, err error) {
	return s.Redis.Lrange(s.WrapKey(key), start, stop)
}

// Lrem is the implementation of redis lrem command.
func (s *XzRedis) Lrem(key string, count int, value string) (val int, err error) {
	return s.Redis.Lrem(s.WrapKey(key), count, value)
}

// Mget is the implementation of redis mget command.
func (s *XzRedis) Mget(keys ...string) (val []string, err error) {
	return s.Redis.Mget(s.WrapKeys(keys...)...)
}

// Persist is the implementation of redis persist command.
func (s *XzRedis) Persist(key string) (val bool, err error) {
	return s.Redis.Persist(s.WrapKey(key))
}

// Pfadd is the implementation of redis pfadd command.
func (s *XzRedis) Pfadd(key string, values ...interface{}) (val bool, err error) {
	return s.Redis.Pfadd(s.WrapKey(key), values...)
}

// Pfcount is the implementation of redis pfcount command.
func (s *XzRedis) Pfcount(key string) (val int64, err error) {
	return s.Redis.Pfcount(s.WrapKey(key))
}

// Pfmerge is the implementation of redis pfmerge command.
func (s *XzRedis) Pfmerge(dest string, keys ...string) error {
	return s.Redis.Pfmerge(dest, s.WrapKeys(keys...)...)
}

// Ping is the implementation of redis ping command.
func (s *XzRedis) Ping() (val bool) {
	return s.Redis.Ping()
}

// Pipelined lets fn to execute pipelined commands.
// fn key must call GetKey or GetKeys to add prefix.
func (s *XzRedis) Pipelined(fn func(gzRedis.Pipeliner) error) (err error) {
	return s.Redis.Pipelined(fn)
}

// Rpop is the implementation of redis rpop command.
func (s *XzRedis) Rpop(key string) (val string, err error) {
	return s.Redis.Rpop(s.WrapKey(key))
}

// Rpush is the implementation of redis rpush command.
func (s *XzRedis) Rpush(key string, values ...interface{}) (val int, err error) {
	return s.Redis.Rpush(s.WrapKey(key), values...)
}

// Sadd is the implementation of redis sadd command.
func (s *XzRedis) Sadd(key string, values ...interface{}) (val int, err error) {
	return s.Redis.Sadd(s.WrapKey(key), values...)
}

// Scan is the implementation of redis scan command.
func (s *XzRedis) Scan(cursor uint64, match string, count int64) (keys []string, cur uint64, err error) {
	return s.Redis.Scan(cursor, match, count)
}

// SetBit is the implementation of redis setbit command.
func (s *XzRedis) SetBit(key string, offset int64, value int) (int, error) {
	return s.Redis.SetBit(s.WrapKey(key), offset, value)
}

// Sscan is the implementation of redis sscan command.
func (s *XzRedis) Sscan(key string, cursor uint64, match string, count int64) (keys []string, cur uint64, err error) {
	return s.Redis.Sscan(s.WrapKey(key), cursor, match, count)
}

// Scard is the implementation of redis scard command.
func (s *XzRedis) Scard(key string) (val int64, err error) {
	return s.Redis.Scard(s.WrapKey(key))
}

// ScriptLoad is the implementation of redis script load command.
func (s *XzRedis) ScriptLoad(script string) (string, error) {
	return s.Redis.ScriptLoad(script)
}

// Set is the implementation of redis set command.
func (s *XzRedis) Set(key, value string) error {
	return s.Redis.Set(s.WrapKey(key), value)
}

// Setex is the implementation of redis setex command.
func (s *XzRedis) Setex(key, value string, seconds int) error {
	return s.Redis.Setex(s.WrapKey(key), value, seconds)
}

// Setnx is the implementation of redis setnx command.
func (s *XzRedis) Setnx(key, value string) (val bool, err error) {
	return s.Redis.Setnx(s.WrapKey(key), value)
}

// SetnxEx is the implementation of redis setnx command with expire.
func (s *XzRedis) SetnxEx(key, value string, seconds int) (val bool, err error) {
	return s.Redis.SetnxEx(s.WrapKey(key), value, seconds)
}

// Sismember is the implementation of redis sismember command.
func (s *XzRedis) Sismember(key string, value interface{}) (val bool, err error) {
	return s.Redis.Sismember(s.WrapKey(key), value)
}

// Smembers is the implementation of redis smembers command.
func (s *XzRedis) Smembers(key string) (val []string, err error) {
	return s.Redis.Smembers(s.WrapKey(key))
}

// Spop is the implementation of redis spop command.
func (s *XzRedis) Spop(key string) (val string, err error) {
	return s.Redis.Spop(s.WrapKey(key))
}

// Srandmember is the implementation of redis srandmember command.
func (s *XzRedis) Srandmember(key string, count int) (val []string, err error) {
	return s.Redis.Srandmember(s.WrapKey(key), count)
}

// Srem is the implementation of redis srem command.
func (s *XzRedis) Srem(key string, values ...interface{}) (val int, err error) {
	return s.Redis.Srem(s.WrapKey(key), values...)
}

// String returns the string representation of s.
func (s *XzRedis) String() string {
	return s.Redis.String()
}

// Sunion is the implementation of redis sunion command.
func (s *XzRedis) Sunion(keys ...string) (val []string, err error) {
	return s.Redis.Sunion(s.WrapKeys(keys...)...)
}

// Sunionstore is the implementation of redis sunionstore command.
func (s *XzRedis) Sunionstore(destination string, keys ...string) (val int, err error) {
	return s.Redis.Sunionstore(destination, s.WrapKeys(keys...)...)
}

// Sdiff is the implementation of redis sdiff command.
func (s *XzRedis) Sdiff(keys ...string) (val []string, err error) {
	return s.Redis.Sdiff(s.WrapKeys(keys...)...)
}

// Sdiffstore is the implementation of redis sdiffstore command.
func (s *XzRedis) Sdiffstore(destination string, keys ...string) (val int, err error) {
	return s.Redis.Sdiffstore(destination, s.WrapKeys(keys...)...)
}

// Sinter is the implementation of redis sinter command.
func (s *XzRedis) Sinter(keys ...string) (val []string, err error) {
	return s.Redis.Sinter(s.WrapKeys(keys...)...)
}

// Sinterstore is the implementation of redis sinterstore command.
func (s *XzRedis) Sinterstore(destination string, keys ...string) (val int, err error) {
	return s.Redis.Sinterstore(destination, s.WrapKeys(keys...)...)
}

// Ttl is the implementation of redis ttl command.
func (s *XzRedis) Ttl(key string) (val int, err error) {
	return s.Redis.Ttl(s.WrapKey(key))
}

// Zadd is the implementation of redis zadd command.
func (s *XzRedis) Zadd(key string, score int64, value string) (val bool, err error) {
	return s.Redis.Zadd(s.WrapKey(key), score, value)
}

// Zadds is the implementation of redis zadds command.
func (s *XzRedis) Zadds(key string, ps ...gzRedis.Pair) (val int64, err error) {
	return s.Redis.Zadds(s.WrapKey(key), ps...)
}

// Zcard is the implementation of redis zcard command.
func (s *XzRedis) Zcard(key string) (val int, err error) {
	return s.Redis.Zcard(s.WrapKey(key))
}

// Zcount is the implementation of redis zcount command.
func (s *XzRedis) Zcount(key string, start, stop int64) (val int, err error) {
	return s.Redis.Zcount(s.WrapKey(key), start, stop)
}

// Zincrby is the implementation of redis zincrby command.
func (s *XzRedis) Zincrby(key string, increment int64, field string) (val int64, err error) {
	return s.Redis.Zincrby(s.WrapKey(key), increment, field)
}

// Zscore is the implementation of redis zscore command.
func (s *XzRedis) Zscore(key, value string) (val int64, err error) {
	return s.Redis.Zscore(s.WrapKey(key), value)
}

// Zrank is the implementation of redis zrank command.
func (s *XzRedis) Zrank(key, field string) (val int64, err error) {
	return s.Redis.Zrank(s.WrapKey(key), field)
}

// Zrem is the implementation of redis zrem command.
func (s *XzRedis) Zrem(key string, values ...interface{}) (val int, err error) {
	return s.Redis.Zrem(s.WrapKey(key), values...)
}

// Zremrangebyscore is the implementation of redis zremrangebyscore command.
func (s *XzRedis) Zremrangebyscore(key string, start, stop int64) (val int, err error) {
	return s.Redis.Zremrangebyscore(s.WrapKey(key), start, stop)
}

// Zremrangebyrank is the implementation of redis zremrangebyrank command.
func (s *XzRedis) Zremrangebyrank(key string, start, stop int64) (val int, err error) {
	return s.Redis.Zremrangebyrank(s.WrapKey(key), start, stop)
}

// Zrange is the implementation of redis zrange command.
func (s *XzRedis) Zrange(key string, start, stop int64) (val []string, err error) {
	return s.Redis.Zrange(s.WrapKey(key), start, stop)
}

// ZrangeWithScores is the implementation of redis zrange command with scores.
func (s *XzRedis) ZrangeWithScores(key string, start, stop int64) (val []gzRedis.Pair, err error) {
	return s.Redis.ZrangeWithScores(s.WrapKey(key), start, stop)
}

// ZRevRangeWithScores is the implementation of redis zrevrange command with scores.
func (s *XzRedis) ZRevRangeWithScores(key string, start, stop int64) (val []gzRedis.Pair, err error) {
	return s.Redis.ZRevRangeWithScores(s.WrapKey(key), start, stop)
}

// ZrangebyscoreWithScores is the implementation of redis zrangebyscore command with scores.
func (s *XzRedis) ZrangebyscoreWithScores(key string, start, stop int64) (val []gzRedis.Pair, err error) {
	return s.Redis.ZrangebyscoreWithScores(s.WrapKey(key), start, stop)
}

// ZrangebyscoreWithScoresAndLimit is the implementation of redis zrangebyscore command with scores and limit.
func (s *XzRedis) ZrangebyscoreWithScoresAndLimit(key string, start, stop int64, page, size int) (val []gzRedis.Pair, err error) {
	return s.Redis.ZrangebyscoreWithScoresAndLimit(s.WrapKey(key), start, stop, page, size)
}

// Zrevrange is the implementation of redis zrevrange command.
func (s *XzRedis) Zrevrange(key string, start, stop int64) (val []string, err error) {
	return s.Redis.Zrevrange(s.WrapKey(key), start, stop)
}

// ZrevrangebyscoreWithScores is the implementation of redis zrevrangebyscore command with scores.
func (s *XzRedis) ZrevrangebyscoreWithScores(key string, start, stop int64) (val []gzRedis.Pair, err error) {
	return s.Redis.ZrevrangebyscoreWithScores(s.WrapKey(key), start, stop)
}

// ZrevrangebyscoreWithScoresAndLimit is the implementation of redis zrevrangebyscore command with scores and limit.
func (s *XzRedis) ZrevrangebyscoreWithScoresAndLimit(key string, start, stop int64, page, size int) (val []gzRedis.Pair, err error) {
	return s.Redis.ZrevrangebyscoreWithScoresAndLimit(s.WrapKey(key), start, stop, page, size)
}

// Zrevrank is the implementation of redis zrevrank command.
func (s *XzRedis) Zrevrank(key, field string) (val int64, err error) {
	return s.Redis.Zrevrank(s.WrapKey(key), field)
}

//// Zunionstore is the implementation of redis zunionstore command.
//func (s *XzRedis) Zunionstore(dest string, store gzRedis.ZStore, keys ...string) (val int64, err error) {
//	return s.Redis.Zunionstore(dest, store, s.WrapKeys(keys...)...)
//}

// Cluster customizes the given Redis as a cluster.
func Cluster() gzRedis.Option {
	return gzRedis.Cluster()
}

// SetSlowThreshold sets the slow threshold.
func SetSlowThreshold(threshold time.Duration) {
	gzRedis.SetSlowThreshold(threshold)
}

// WithPass customizes the given Redis with given password.
func WithPass(pass string) gzRedis.Option {
	return gzRedis.WithPass(pass)
}

// WithTLS customizes the given Redis with TLS enabled.
func WithTLS() gzRedis.Option {
	return gzRedis.WithTLS()
}
