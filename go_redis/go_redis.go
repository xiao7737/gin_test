package go_redis

import (
	"gin_test/conf"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Redis struct {
	RedisConn *redis.ClusterClient
}

var RedisCluster *Redis

func init_() {
	RedisCluster = &Redis{
		RedisConn: GetConnect(),
	}
}

func GetConnect() *redis.ClusterClient {
	config := conf.LoadConfig()

	//connect to redis-cluster
	cluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{"127.0.0.1:7001", "127.0.0.1:7002"},
		Password: config.RedisCluster.Password,
		/*ReadOnly:      true,
		RouteRandomly: true, //随机选择redis节点
		PoolSize:      config.RedisCluster.MaxPoolSize,
		DialTimeout:   5 * time.Second,
		MaxConnAge:    10 * time.Second,
		IdleTimeout:   60 * time.Second,*/
	})

	// ping redis-cluster
	_, err := cluster.Ping().Result()
	if err != nil {
		log.Fatalf("redis-cluster connect err: %v", err)
	} else {
		log.Println("redis-cluster connect success")
	}
	return cluster
}

// get a value for string
func Get(key string) (string, error) {
	client := RedisCluster.RedisConn
	defer client.Close()
	value, err := client.Get(key).Result()
	if err == redis.Nil {
		return "", err
	} else if err != nil {
		return "", err
	} else {
		return value, nil
	}
}

func Set(key string, value interface{}, time time.Duration) error {
	client := RedisCluster.RedisConn
	err := client.Set(key, value, time).Err()
	if err != nil {
		return err
	}
	return nil
}
