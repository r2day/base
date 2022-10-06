package middle

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v9"
	"time"

	"github.com/r2day/base/log"
	btime "github.com/r2day/base/time"
	"github.com/r2day/enum"
)

// CacheClient 缓存
type CacheClient struct {
	Ctx  context.Context
	Conn *redis.Client
}

// InitCache 初始化redis
func InitCache(ctx context.Context, dsn string, user string, password string, readTimeout time.Duration) CacheClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:        dsn,
		Username:    user,
		Password:    password,        // no password set
		DB:          defaultDatabase, // use default DB
		ReadTimeout: readTimeout,
	})
	rc := CacheClient{
		Conn: rdb,
		Ctx:  ctx,
	}
	return rc
}

// SetSmsCode 设置短信验证码
func (rc *CacheClient) SetSmsCode(phone string, code string) {

	// 每天会使用日期作为key的一部分
	finalKey := fmt.Sprintf("%s_%s_%s", enum.SmsCode, btime.GetDaily(), phone)
	rc.Conn.SetEx(rc.Ctx, finalKey, code, enum.SmsCodeExpireTime) // 验证码有效期是5分钟 (1分钟后重发的验证码会覆盖旧的)

}

// CheckSmsCode 校验验证码是否正确
func (rc *CacheClient) CheckSmsCode(phone string, code string) bool {
	// 每天会使用日期作为key的一部分
	finalKey := fmt.Sprintf("%s_%s_%s", enum.SmsCode, btime.GetDaily(), phone)
	val, err := rc.Conn.Get(rc.Ctx, finalKey).Result()
	if err != nil {
		// panic(err)
		log.Logger.Warn("no data for read")
		log.Logger.WithField("finalyKey", finalKey).Info("but keep next")
		return false
	}
	if val != code {
		return false
	}
	// 验证通过后自动删除
	rc.Conn.Del(rc.Ctx, finalKey)
	return true
}
