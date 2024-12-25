package cache

import (
	"context"
	"gin-vben-admin/global"
	"github.com/allegro/bigcache/v3"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/marshaler"
	"github.com/eko/gocache/lib/v4/store"
	bigcache_store "github.com/eko/gocache/store/bigcache/v4"
	redis_store "github.com/eko/gocache/store/redis/v4"
	"github.com/redis/go-redis/v9"
	"time"
)

func Init() {
	bigcacheClient, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(30*time.Minute))
	bigcacheStore := bigcache_store.NewBigcache(bigcacheClient)

	global.BigCache = marshaler.New(cache.New[any](bigcacheStore))

	redisStore := redis_store.NewRedis(redis.NewClient(&redis.Options{
		Addr:     global.Conf.Redis.Network,
		Password: global.Conf.Redis.Password,
		DB:       global.Conf.Redis.DB,
	}))
	global.RedisCache = marshaler.New(cache.New[any](redisStore))
	redisTStore := redis_store.NewRedis(redis.NewClient(&redis.Options{
		Addr:     global.Conf.Redis.Network,
		Password: global.Conf.Redis.Password,
		DB:       global.Conf.Redis.TokenDB,
	}))

	global.TokenCache = marshaler.New(cache.New[any](redisTStore))
}

func Remember(cache *marshaler.Marshaler, key string, call func() (any, error), expires time.Duration, returnObj any) (interface{}, error) {
	ctx := context.Background()
	_, err := cache.Get(ctx, key, returnObj)
	if err == nil && returnObj != nil {
		return returnObj, nil
	}
	data, err := call()
	if err != nil {
		return nil, err
	}
	if data != nil {
		cache.Set(ctx, key, data, store.WithExpiration(expires))
		cache.Get(ctx, key, returnObj) // 指针变量不可直接复制
	}
	return data, nil

}
