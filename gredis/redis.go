package gredis

import (
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

func Get(key string) ( /*[]byte*/ string, error) {
	conn := RedisConn.Get() //连接池中获取一个活跃连接
	defer conn.Close()
	if reply, err := redis.String(conn.Do("GET", key)); err != nil {
		return "", err
	} else {
		return reply, nil
	}

	/*conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil*/
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

func Set(key string, value interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()
	/*data, err := json.Marshal(data)
	if err != nil {
		return err
	}*/
	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("DEL", key))
}
