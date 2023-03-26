package redis

import (
	"blog/db"
	"blog/logger"
	"context"
	"time"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type (
	Cache[T Type] interface {
		SetDefault(key string, value T)
		Set(key string, value T, expiration time.Duration)
		Get(key string) (T, bool)
		Delete(key string)
		BatchDelete(keys []string)
	}
	Type interface {
		// entry.Article | []entry.Article | entry.Article_tags
		any
	}
)

var client *redis.Client
var ctx context.Context

func init() {
	client = db.GetRedisClient()
	ctx = context.Background()
}

func Get[T Type](key string) *T {
	result := client.Get(ctx, key)
	if errors.Is(result.Err(), redis.Nil) {
		logger.Log.Infof("查询的key：%s 不存在", key)
		return nil
	}
	var res T
	err := json.Unmarshal([]byte(result.Val()), &res)
	if err != nil {
		logger.Log.Error(err)
		return nil
	}
	return &res

}
func Set[T Type](key string, val T, expirationTime ...Option) string {
	marshal, err := json.Marshal(val)
	if err != nil {
		return "convert fail"
	}
	expiration := 0 * time.Second
	if len(expirationTime) != 0 {
		expiration = expirationTime[0]()
	}
	set := client.Set(ctx, key, marshal, expiration)
	if set.Err() != nil {
		logger.Log.Error(set.Err())
		return "Fail"
	}
	return set.Val()
}

type Option func() time.Duration

func WithExpiration(expiration time.Duration) Option {
	return func() time.Duration {
		return expiration
	}
}
