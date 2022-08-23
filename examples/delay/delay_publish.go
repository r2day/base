package main

import (
        "log"
        "os"
        "strings"

        amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
        if err != nil {
                log.Panicf("%s: %s", msg, err)
        }
}

func main() {
        conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
        failOnError(err, "Failed to connect to RabbitMQ")
        defer conn.Close()

        ch, err := conn.Channel()
        failOnError(err, "Failed to open a channel")
        defer ch.Close()

	// .exchangeDeclare("my-exchange", "x-delayed-message", true, false, args)
        args := make(amqp.Table)
	args["x-delayed-type"] = "direct"

	 //err = ch.ExchangeDeclare(
	   //ch.Exchange, // name
	 // "x-delayed-message",      // type
	  //true,          // durable
	  //false,         // auto-deleted
	  //false,         // internal
	  //false,         // no-wait
	  //args,           // arguments
	//)
        
	//err = ch.ExchangeDeclare("delayed", "x-delayed-message", true, false, false, false, args)
	err = ch.ExchangeDeclare("direct2", "direct", true, false, false, false, nil)
        failOnError(err, "Failed to declare a queue")

        q, err := ch.QueueDeclare(
                "task_queue3", // name
                true,         // durable
                false,        // delete when unused
                false,        // exclusive
                false,        // no-wait
                nil,          // arguments
        )
        failOnError(err, "Failed to declare a queue")

        body := bodyFrom(os.Args)
        err = ch.Publish(
                "direct2",           // exchange
                q.Name,       // routing key
                false,        // mandatory
                false,
                amqp.Publishing{
                        DeliveryMode: amqp.Persistent,
                        ContentType:  "text/plain",
                        Body:         []byte(body),
                       // Headers:     amqp.Table{"x-delay": 500},
                })
        failOnError(err, "Failed to publish a message")
        log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
        var s string
        if (len(args) < 2) || os.Args[1] == "" {
                s = "hello"
        } else {
                s = strings.Join(args[1:], " ")
        }
        return s
}
