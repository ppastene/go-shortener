package cache

import (
	"errors"
	"sync"
	"time"

	"github.com/ppastene/go-shortener/internal/config"
	"github.com/ppastene/go-shortener/internal/domain"
)

type MemoryCache struct {
	data   map[string]domain.Redirect
	mu     sync.RWMutex
	config *config.Config
}

func NewMemoryCache(cfg *config.Config) *MemoryCache {
	cache := &MemoryCache{
		data:   make(map[string]domain.Redirect),
		config: cfg,
	}
	go cache.startCleanUp()
	return cache
}

func (m *MemoryCache) Get(key string) (domain.Redirect, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	item, found := m.data[key]
	if !found {
		return domain.Redirect{}, errors.New("Key not found")
	}
	if time.Now().After(item.Expiration) {
		m.Delete(key)
		return domain.Redirect{}, errors.New("Key expired")
	}
	return item, nil
}

func (m *MemoryCache) Set(key string, shortenedUrl domain.Redirect) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = domain.Redirect{
		Url:        shortenedUrl.Url,
		Expiration: shortenedUrl.Expiration,
	}
}

func (m *MemoryCache) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func (m *MemoryCache) List() map[string]domain.Redirect {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.data
}

func (m *MemoryCache) startCleanUp() {
	for {
		time.Sleep(15 * time.Second)
		m.cleanUp()
	}
}

func (m *MemoryCache) cleanUp() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for key, item := range m.data {
		if time.Now().After(item.Expiration) {
			delete(m.data, key)
		}
	}
}
