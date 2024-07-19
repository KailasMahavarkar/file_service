package cache

import (
	"time"
)

func NewURLCache() *URLCache {
	return &URLCache{
		cache: make(map[string]CacheEntry),
	}
}

func (c *URLCache) Get(key string) (string, bool) {
	c.mutex.RLock()
	entry, found := c.cache[key]
	c.mutex.RUnlock()

	if found && time.Now().Before(entry.ExpiryTime) {
		return entry.URL, true
	}

	return "", false
}

func (c *URLCache) Set(key string, url string, expiry time.Time) {
	c.mutex.Lock()
	c.cache[key] = CacheEntry{
		URL:        url,
		ExpiryTime: expiry,
	}

	c.mutex.Unlock()
}

// run a cron job to clear the cache every 5 minutes
func (c *URLCache) Clear() {
	c.mutex.Lock()
	for key, entry := range c.cache {
		if time.Now().After(entry.ExpiryTime) {
			delete(c.cache, key)
		}
	}
	c.mutex.Unlock()
}

// Path: pkg\cache\types.go
