package conn

import (
	"gin_test/conf"
	. "gin_test/msg"
	"github.com/streadway/amqp"
)

func GetRabbitAndQueue() (ch *amqp.Channel) {
	cf := conf.LoadConfig()
	conn, err := amqp.Dial("amqp://" + cf.RabbitMQ.User + ":" + cf.RabbitMQ.Password + "@" + cf.RabbitMQ.Address + "/")
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err = conn.Channel()
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
	return ch
}
