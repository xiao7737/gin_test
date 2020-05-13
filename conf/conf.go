package conf

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type App struct {
	RunPort   string
	PProfPort string
}

type Database struct {
	Driver   string
	Address  string
	Database string
	User     string
	Password string
}
type Redis struct {
	Address  string
	Password string
}
type RedisCluster struct {
	Password    string
	MaxPoolSize int
}
type RabbitMQ struct {
	Address  string
	User     string
	Password string
}
type MongoDB struct {
	Address     string
	MaxPoolSize uint64
}

type Configuration struct {
	App          App
	Db           Database
	Redis        Redis
	RedisCluster RedisCluster //go-redis
	RabbitMQ     RabbitMQ
	MongoDB      MongoDB
}

var config *Configuration
var once sync.Once

func LoadConfig() *Configuration {
	once.Do(func() {
		// Open("config.json")  简约版
		file, err := os.OpenFile("config.json", os.O_RDONLY, 400) //只读模式打开
		if err != nil {
			log.Fatalln("cannot open config file", err)
		}
		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err != nil {
			log.Fatalln("cannot get config from file", err)
		}
	})
	return config
}
