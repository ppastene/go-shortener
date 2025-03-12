package cache

import (
	"github.com/ppastene/go-shortener/internal/domain"
)

type Cache interface {
	Get(key string) (domain.Rediect, error)
	Set(key string, shortenerUrl domain.Rediect)
	Delete(key string)
	List() map[string]domain.Rediect
}
