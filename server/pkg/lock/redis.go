package lock

import (
	"context"
	"fmt"
	"gin-vben-admin/global"
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

var locker *redislock.Client

func Init() {
	client := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     global.Conf.Redis.Network,
		Password: global.Conf.Redis.Password,
		DB:       global.Conf.Redis.DB,
	})
	locker = redislock.New(client)
}

func Lock(key string, ttl time.Duration, proc bool) (*redislock.Lock, context.Context, error) {
	if proc {
		key = fmt.Sprintf("%s_%d", key, os.Getpid())
	}
	ctx := context.Background()
	lock, err := locker.Obtain(ctx, "_lock_"+key, ttl, nil)
	if err != nil {
		return nil, ctx, err
	}
	return lock, ctx, nil
}

func LockOpt(key string, ttl time.Duration, duration time.Duration, maxtime int, proc bool) (*redislock.Lock, context.Context, error) {
	if proc {
		key = fmt.Sprintf("%s_%d", key, os.Getpid())
	}
	ctx := context.Background()
	lock, err := locker.Obtain(ctx, "_lock_"+key, ttl, &redislock.Options{
		RetryStrategy: redislock.LimitRetry(redislock.LinearBackoff(duration), maxtime),
	})
	if err != nil {
		return nil, ctx, err
	}
	return lock, ctx, nil
}
