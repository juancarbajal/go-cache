package cache

import (
	"context"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

type Object struct {
	Str string
}
type TCacheRedisOptions struct {
	host     string
	port     string
	password string
}
type TCacheRedis struct {
	ICacheManager
	c *cache.Cache
}

// (cr *TCacheRedis) Create ...
func (cr TCacheRedis) Init(host string, port string, password string) {
	ropt := redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	}
	rdb := redis.NewClient(&ropt)
	cr.c = cache.New(&cache.Options{
		Redis: rdb,
	})
}

func (cr TCacheRedis) Add(key string, value string, expiration uint32) error {
	obj := &Object{
		Str: value,
	}
	if err := cr.c.Set(&cache.Item{
		Ctx:   context.TODO(),
		Key:   key,
		Value: obj,
		TTL:   time.Duration(expiration) * time.Second,
	}); err != nil {
		return err
	}
	return nil
}
func (cr TCacheRedis) Remove(key string) {
	cr.c.Delete(context.TODO(), key)
}
func (cr TCacheRedis) Find(key string) (string, error) {
	var wanted string
	if err := cr.c.Get(context.TODO(), key, &wanted); err == nil {
		return "", err
	}
	return wanted, nil
}
