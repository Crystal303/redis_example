package redis_example

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	addr     = "127.0.0.1:6379"
	db       = 1
	password = ""
)

var RedisPool = newPool(addr, password, db)

func newPool(addr, password string, db int) *redis.Pool {
	return &redis.Pool{
		//DialContext: func(ctx context.Context) (redis.Conn, error) {
		//	c, err := redis.DialContext(ctx, "tcp", addr)
		//	if err != nil {
		//		return nil, err
		//	}
		//	if password != "" {
		//		if _, err := c.Do("AUTH", password); err != nil {
		//			_ = c.Close()
		//			return nil, err
		//		}
		//	}
		//	if _, err := c.Do("SELECT", db); err != nil {
		//		_ = c.Close()
		//		return nil, err
		//	}
		//	return c, nil
		//},
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", db); err != nil {
				_ = c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         3,
		MaxActive:       0,
		IdleTimeout:     240 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}
