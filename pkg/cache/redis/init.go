package redis

import (
	"context"
	"fmt"
	"time"
	"watt/pkg/utils"

	"github.com/go-redis/redis/v8"
)

var (
	redis_link *redis.Client
	err        error
	ctx        = context.Background()
)

func init() {

	redis_link = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", utils.Setting.Database.Redis.Host, utils.Setting.Database.Redis.Port),
		Password:    utils.Setting.Database.Redis.Password,
		DB:          utils.Setting.Database.Redis.Database,
		DialTimeout: utils.Setting.Database.Redis.Timeout * time.Second,
	})

	_, err = redis_link.Ping(ctx).Result()

	if err != nil {
		utils.Exit(err)
	}
}
