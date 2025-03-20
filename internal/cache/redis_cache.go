package cache

import (
	"github.com/ppastene/go-shortener/internal/config"
	"github.com/ppastene/go-shortener/internal/domain"
)

type RedisCache struct {
	Config *config.Config
}

func NewRedisCache(cfg *config.Config) *RedisCache {
	return &RedisCache{cfg}
}

func (r *RedisCache) Get(key string) (domain.Redirect, error) {
	return domain.Redirect{}, nil
}
func (r *RedisCache) Set(key string, shortenerUrl domain.Redirect) {}
func (r *RedisCache) Delete(key string)                            {}
func (r *RedisCache) List() map[string]domain.Redirect {
	return map[string]domain.Redirect{}
}
