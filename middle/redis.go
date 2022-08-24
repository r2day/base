package middle

import (
	"context"
	"fmt"
	"strconv"
	"time"

	redis "github.com/go-redis/redis/v9"

	btime "github.com/r2day/base/time"
	"github.com/r2day/enum"
)

const (
	dailyExpireTime = 24 * 60 * 60 * time.Second

	defaultDatabase = 0
)

type RedisClient struct {
	Ctx  context.Context
	Conn *redis.Client
}

// InitRedis 初始化redis
func InitRedis(ctx context.Context, dsn string, user string, password string) RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: password,        // no password set
		DB:       defaultDatabase, // use default DB
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
		panic(err)
	}

	// 未初始化
	if val == "" {
		newVal := begin + 1
		// SET key value EX 10 NX
		_, err = rc.Conn.SetEx(rc.Ctx, finalKey, newVal, dailyExpireTime).Result()
		if err != nil {
			panic(err)
		}
		val = strconv.Itoa(newVal)
	}

	_, err = rc.Conn.Incr(rc.Ctx, finalKey).Result()
	if err != nil {
		panic(err)
	}

	_, err = rc.Conn.ExpireNX(rc.Ctx, finalKey, dailyExpireTime).Result()
	if err != nil {
		panic(err)
	}

	// SET key value EX 10 NX
	// set, err := rdb.SetEX(rc.Ctx, finalKey, newVal, dailyExpireTime).Result()
	finalVal := fmt.Sprintf("%s%s", prefix, val)
	return finalVal

}
