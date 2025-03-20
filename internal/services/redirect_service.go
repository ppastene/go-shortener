package services

import (
	"errors"
	"time"

	"github.com/ppastene/go-shortener/internal/cache"
	"github.com/ppastene/go-shortener/internal/config"
	"github.com/ppastene/go-shortener/internal/domain"
)

type ShortenerService struct {
	cache cache.Cache
	cfg   *config.Config
}

func NewShortenerService(cache cache.Cache, cfg config.Config) *ShortenerService {
	return &ShortenerService{
		cache: cache,
		cfg:   &cfg,
	}
}

func (se ShortenerService) GetUrl(key string) (domain.Redirect, error) {
	shortenedUrl, err := se.cache.Get(key)
	if err != nil {
		return domain.Redirect{}, err
	}
	return shortenedUrl, nil
}

func (se ShortenerService) SaveUrl(key, url string) error {
	if se.isShortcodeExists(key) {
		return errors.New("Key already exists")
	}
	shortenedUrl := domain.Redirect{
		Url:        url,
		Expiration: time.Now().Add(time.Second * time.Duration(se.cfg.ExpirationTime)),
	}
	se.cache.Set(key, shortenedUrl)
	return nil
}

func (se ShortenerService) ListShortcodes() map[string]domain.Redirect {
	return se.cache.List()
}

func (se ShortenerService) isShortcodeExists(key string) bool {
	shortenedUrl, _ := se.cache.Get(key)
	if (shortenedUrl == domain.Redirect{}) {
		return false
	}
	return true
}
