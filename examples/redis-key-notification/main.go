package main

import (
	//"github.com/go-redis/redis"
	"context"
	"fmt"
	"log"
	"time"
	"unsafe"

	redis "github.com/go-redis/redis/v9"
)

type PSubscribeCallback func(pattern, channel, message string)

type PSubscriber struct {
	client redis.PubSub
	cbMap  map[string]PSubscribeCallback
}

func (c *PSubscriber) PConnect(ctx context.Context, ip string, port uint16) {
	// conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	// if err != nil {
	// 	log.Critical("redis dial failed.")
	// }
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		// Password: password,        // no password set
		// DB:       defaultDatabase, // use default DB
		ReadTimeout: -1,
	})
	c.client = *rdb.PSubscribe(ctx)
	c.cbMap = make(map[string]PSubscribeCallback)

	go func() {
		for {
			log.Println("wait...")
			res, _ := c.client.Receive(context.TODO())
			switch s := res.(type) {
			case *redis.Message:
				fmt.Printf("got a event, call your callback now", )
				pattern := (*string)(unsafe.Pointer(&s.Pattern))
				channel := (*string)(unsafe.Pointer(&s.Channel))
				message := (*string)(unsafe.Pointer(&s.Payload))

				fmt.Printf("target channel name: [%v]\n", channel)
				c.cbMap[*channel](*pattern, *channel, *message)
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", s.Channel, s.Kind, s.Count)
			case error:
				log.Fatal("error handle...")
				continue
			default:
				fmt.Printf("nothing match: event=[%v]", s)
			}
			
		}
	}()

}
func (c *PSubscriber) Psubscribe(ctx context.Context, patterns string, cb PSubscribeCallback) {
	err := c.client.PSubscribe(ctx, patterns)
	if err != nil {
		log.Fatal("redis Subscribe error.")
	}

	c.cbMap[patterns] = cb
}

func TestPubCallback(patter, chann, msg string) {
	log.Println("TestPubCallback patter : "+patter+" channel : ", chann, " message : ", msg)
}

func main() {

	log.Println("===========main start============")
	ctx := context.TODO()
	var psub PSubscriber
	psub.PConnect(ctx, "127.0.0.1", 6397)
	psub.Psubscribe(ctx, "__keyevent@0__:expired", TestPubCallback)
	//It can also be: `__keyspace@0__:cool`
	for {
		time.Sleep(1 * time.Second)
	}
}
