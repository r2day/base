package middle

import (
	amqp "github.com/rabbitmq/amqp091-go"

	logger "github.com/r2day/base/log"
)

func failOnError(err error, msg string) {
	if err != nil {
		logger.Logger.Panicf("%s: %s", msg, err)
	}
}

type Amqp struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func InitAMQP(address string, name string) Amqp {

	var instance Amqp
	conn, err := amqp.Dial(address)
	failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()

	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")
	instance.Conn = conn
	instance.Channel = ch
	instance.Queue = q

	return instance

}

// Close 关闭
func (i Amqp) Close() {
	i.Conn.Close()
	i.Channel.Close()
}

func (i Amqp) Send(payload []byte) {

	err := i.Channel.Publish(
		"",           // exchange
		i.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        payload,
		})
	failOnError(err, "Failed to publish a message")
}

func (i Amqp) Reiceive(f func([]byte) []byte) {
	msgs, err := i.Channel.Consume(
		i.Queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			logger.Logger.Debugf("Received a message: %s", d.Body)
			f(d.Body)
		}
	}()

	logger.Logger.Warn(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
