package main

import (
	"fmt"
	"gin_test/conf"
	. "gin_test/middleware/rabbitmq/conn"
	"github.com/streadway/amqp"
	"math/rand"
	"strconv"
)

func main() {
	cf := conf.LoadConfig()
	conn, err := amqp.Dial("amqp://" + cf.RabbitMQ.User + ":" + cf.RabbitMQ.Password + "@" + cf.RabbitMQ.Address + "/")
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declaring a queue is idempotent - queue will only be created if it doesn't exist already.
	_, err = ch.QueueDeclare(
		"gin_test_queue", //queue name
		false,            //durable
		false,            //autoDelete
		false,            //exclusive
		false,            //no-wait
		map[string]interface{}{ //arguments map
			"x-queue-type": "classic", //type of queue
		},
	)
	FailOnError(err, "Failed to declare a queue")

	for i := 0; i < 100; i++ { //测试发100次
		msg := "hello rabbit " + strconv.Itoa(rand.Int())
		err = ch.Publish(
			"gin_test",          //exchange name
			"gin_routing_value", //routing key
			false,               //mandatory
			false,               //immediate
			amqp.Publishing{
				ContentType:  "text/plain",
				Body:         []byte(msg),
				Priority:     9, //优先级，范围0-9  rabbit绑定exchange设置参数"x-max-priority": 9
				DeliveryMode: 1, //2是持久化，消息的持久化需要队列持久化支持
				// rabbit为啥不用bool来表示消息是否持久化，而是unit的1和2来表示
				//2020-04-15 明白了，源码介绍目前是两个枚举，还有其它模式，用枚举方便扩展！
			})
		if err != nil {
			FailOnError(err, "Failed to publish a message")
		}
	}
	fmt.Println("Send success!")
}
