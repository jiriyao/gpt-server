package redis

//import (
//	"testing"
//
//	"github.com/gogf/gf/frame/g"
//)
//
//func TestScript_Do(t *testing.T) {
//	////eval "return redis.call('get', 'foo')" 0
//	do, err := g.Redis("config").Do("eval", "return redis.call('get', 'foo')", 0)
//	t.Log(do, err)
//
//	lua := `
//	return redis.call('get', 'foo')
//	`
//	script := NewScript(0, lua)
//	r, err := script.Do(g.Redis("config"))
//	t.Log(string(r.([]uint8)), err)
//}
