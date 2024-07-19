package cache

import (
	"sync"
	"time"
)

type URLCache struct {
	cache map[string]CacheEntry
	mutex sync.RWMutex
}

type CacheEntry struct {
	URL        string
	ExpiryTime time.Time
}
