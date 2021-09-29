package redlock

import (
	"redis_example"

	"github.com/gomodule/redigo/redis"
)

const scriptCAS = `
	if redis.call("get",KEYS[1]) == ARGV[1] then
		return redis.call("del",KEYS[1])
	else
		return 0
	end
`

const DefaultTimeout = 5

func (m Mutex) Unlock() (int, error) {
	lua := redis.NewScript(1, scriptCAS)

	return redis.Int(lua.Do(m.conn, m.key, m.value))
}

type Mutex struct {
	key     string
	value   string
	conn    redis.Conn
	timeout int
}

func NewMutex(key, value string, timeout int) Mutex {
	if timeout == 0 {
		timeout = DefaultTimeout
	}
	return Mutex{
		key:     "lock:" + key,
		value:   value,
		conn:    redis_example.RedisPool.Get(),
		timeout: timeout,
	}
}

func (m Mutex) Lock() (bool, error) {
	_, err := redis.String(m.conn.Do("SET", m.key, m.value, "EX", m.timeout, "NX"))
	if err != nil {
		// 已存在键值
		if err == redis.ErrNil {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
