package gredis

import (
	"gin_test/conf"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConn *redis.Pool

func init() {
	config := conf.LoadConfig() //虽然此处调用了loadConfig，由于是单例，初始化server就已经获取配置，这里会直接返回实例
	RedisConn = &redis.Pool{
		MaxIdle:         30, //最大空闲连接
		MaxActive:       30, //最大连接数，0没有限制
		IdleTimeout:     200,
		Wait:            false,
		MaxConnLifetime: 0,
		Dial: func() (conn redis.Conn, err error) { //配置连接
			c, err := redis.Dial("tcp", config.Redis.Address)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", config.Redis.Password); err != nil {
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
