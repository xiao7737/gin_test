package gredis

import (
	"encoding/json"
	"gin_test/conf"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:         30, //最大空闲连接
		MaxActive:       30, //最大连接数，0没有限制
		IdleTimeout:     200,
		Wait:            false,
		MaxConnLifetime: 0,
		Dial: func() (conn redis.Conn, err error) { //配置连接
			c, err := redis.Dial("tcp", conf.HOST)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", conf.AUTH); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { //检测活度
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get() //连接池中获取一个活跃连接
	defer conn.Close()
	if reply, err := redis.Bytes(conn.Do("GET", key)); err != nil {
		return nil, err
	} else {
		return reply, nil
	}
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	if exists, err := redis.Bool(conn.Do("EXISTS", key)); err != nil {
		return false
	} else {
		return exists
	}
}

func Set(key string, data interface{}, time int) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}
	reply, err := redis.Bool(conn.Do("SET", key, value))
	_, _ = conn.Do("EXPIRE", key, time)
	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("DEL", key))
}
