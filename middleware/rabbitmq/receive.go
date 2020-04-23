package main

import (
	"gin_test/conf"
	. "gin_test/msg"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	cf := conf.LoadConfig()
	conn, err := amqp.Dial("amqp://" + cf.RabbitMQ.User + ":" + cf.RabbitMQ.Password + "@" + cf.RabbitMQ.Address + "/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"gin_test_queue", //queue name
		true,             //durable
		false,            //autoDelete
		false,            //exclusive
		false,            //no-wait
		map[string]interface{}{
			"x-queue-type": "classic",
		},
	)
	FailOnError(err, "Failed to declare a queue")
	messages, err := ch.Consume(
		q.Name, //queue name
		"",     //consumer
		true,   //auto-ack
		false,  //exclusive
		false,  //un-local
		false,  //no-wait
		nil,    //args
	)
	FailOnError(err, "Failed to register a consumer")

	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	//todo 将消息队列的消息放到slice或者channel
	//log.Fatalf("Waiting for messages.  press CTRL+C to exit")
}
