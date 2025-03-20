package cache

import (
	"github.com/ppastene/go-shortener/internal/domain"
)

type Cache interface {
	Get(key string) (domain.Redirect, error)
	Set(key string, shortenerUrl domain.Redirect)
	Delete(key string)
	List() map[string]domain.Redirect
}
