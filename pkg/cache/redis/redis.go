package redis

import "time"

type rediscli struct {
}

func newRedis() *rediscli {
	return new(rediscli)
}

var Cli = newRedis()

// 设置key
func (r *rediscli) Set(key, val string, expire int) error {

	return redis_link.Set(ctx, key, val, time.Duration(expire)*time.Second).Err()
}

// 获取key
func (r *rediscli) Get(key string) (string, error) {

	return redis_link.Get(ctx, key).Result()
}

// 删除key
func (r *rediscli) Del(key string) error {

	return redis_link.Del(ctx, key).Err()
}
