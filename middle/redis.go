package middle

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"unsafe"

	redis "github.com/go-redis/redis/v9"

	"github.com/r2day/base/log"
	btime "github.com/r2day/base/time"
	"github.com/r2day/enum"
)

const (
	dailyExpireTime = 24 * 60 * 60 * time.Second

	defaultDatabase = 0
)

type PSubscribeCallback func(pattern, channel, message string)

type RedisClient struct {
	Ctx        context.Context
	Conn       *redis.Client
	PSubClient redis.PubSub
	cbMap      map[string]PSubscribeCallback
}

// InitRedis 初始化redis
func InitRedis(ctx context.Context, dsn string, user string, password string, readTimeout time.Duration) RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: password,        // no password set
		DB:       defaultDatabase, // use default DB
		ReadTimeout: readTimeout,
	})
	rc := RedisClient{
		Conn: rdb,
		Ctx:  ctx,
	}
	return rc
}

// GetStoreQueueSeq 获取当天排队号
// key redis中的key
// prefix 生成的序列号的前缀
// length 排队号长度
// 例如: set("mykey", "R20001")
func (rc *RedisClient) GetStoreQueueSeq(key string, prefix string, begin int) string {

	// 每天会使用日期作为key的一部分
	finalKey := fmt.Sprintf("%s_%s_%s_%s", enum.QueueSeq, btime.GetDaily(), key, prefix)
	val, err := rc.Conn.Get(rc.Ctx, finalKey).Result()
	if err != nil {
		// panic(err)
		log.Logger.Warn("no data for read")
		log.Logger.WithField("finalyKey", finalKey).Info("but keep next")
	}

	// 未初始化
	if val == "" {
		newVal := begin + 1
		// SET key value EX 10 NX
		_, err = rc.Conn.SetEx(rc.Ctx, finalKey, newVal, dailyExpireTime).Result()
		if err != nil {
			// panic(err)
			log.Logger.Error(err)
			log.Logger.WithField("finalyKey", finalKey).Info("set key failed")
		}
		val = strconv.Itoa(newVal)
	}

	_, err = rc.Conn.Incr(rc.Ctx, finalKey).Result()
	if err != nil {
		log.Logger.Error(err)
		log.Logger.WithField("finalyKey", finalKey).Info("set key failed")
	}

	_, err = rc.Conn.ExpireNX(rc.Ctx, finalKey, dailyExpireTime).Result()
	if err != nil {
		log.Logger.Error(err)
		log.Logger.WithField("finalyKey", finalKey).Info("set key failed")
	}

	// SET key value EX 10 NX
	// set, err := rdb.SetEX(rc.Ctx, finalKey, newVal, dailyExpireTime).Result()
	finalVal := fmt.Sprintf("%s%s", prefix, val)
	return finalVal

}

// SetOrderPlace 下单后设置订单id缓存
// 目的有两个
// 1. 当支付的时候会去检查是否还存在值，如果已经过期无法获取到该key就不再允许支付
// 2. 当过期后会监听事件(flow-consume会处理) 例如: 关单、回退库存
func (rc *RedisClient) SetOrderPlace(ctx context.Context, orderId string) {

	orderPlaceKey := fmt.Sprintf("%s_%s", enum.OrderPlaceTimeTTL, orderId)
	minutes := btime.SecondsToMinutes(int(enum.DefaultOrderExpireTime.Seconds()))
	rc.Conn.SetEx(ctx, orderPlaceKey, minutes, enum.DefaultOrderExpireTime)

}

// GetOrderPlace 检查订单是否还可以支付
func (rc *RedisClient) GetOrderPlace(ctx context.Context, orderId string) string {

	finalKey := fmt.Sprintf("%s_%s", enum.OrderPlaceTimeTTL, orderId)
	val, err := rc.Conn.Get(rc.Ctx, finalKey).Result()
	if err != nil {
		log.Logger.Warn("no data for read")
		log.Logger.WithField("finalyKey", finalKey).Info("but keep next")
		return ""
	}
	return val
}

// DelOrderPlace 当前订单已经支付完成，即可以删除该key
// 避免重复支付
func (rc *RedisClient) DelOrderPlace(ctx context.Context, orderId string) {
	finalKey := fmt.Sprintf("%s_%s", enum.OrderPlaceTimeTTL, orderId)
	rc.Conn.Del(rc.Ctx, finalKey)
}

func (rc *RedisClient) Psubscribe(ctx context.Context, patterns string, cb PSubscribeCallback) {
	err := rc.PSubClient.PSubscribe(ctx, patterns)
	if err != nil {
		log.Logger.Fatal(err)
	}

	rc.cbMap[patterns] = cb
}

// ListenEvent 监听事件
// 避免重复支付
func (rc *RedisClient) ListenEvent(ctx context.Context) {
	rc.PSubClient = *rc.Conn.PSubscribe(ctx)
	rc.cbMap = make(map[string]PSubscribeCallback)

	go func() {
		for {
			log.Logger.Info("waitting for event")
			res, _ := rc.PSubClient.Receive(context.TODO())
			switch s := res.(type) {
			case *redis.Message:
				fmt.Printf("got a event, call your callback now")
				pattern := (*string)(unsafe.Pointer(&s.Pattern))
				channel := (*string)(unsafe.Pointer(&s.Channel))
				message := (*string)(unsafe.Pointer(&s.Payload))

				fmt.Printf("target channel name: [%v]\n", channel)
				rc.cbMap[*channel](*pattern, *channel, *message)
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", s.Channel, s.Kind, s.Count)
			case error:
				log.Logger.Fatal("error handler")
				continue
			default:
				log.Logger.Error("the messsage is not matched")
			}

		}
	}()
}
