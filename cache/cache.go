package cache

import (
	"errors"
	"io"
	"time"
)

var (
	ErrCacheMiss          = errors.New("cache miss")
	ErrCacheEntryNotFound = errors.New("cache entry not found")
)

type Cache[ObjectData any] interface {
	// Get retrieves an entry from the cache by its input key.
	Get(key CacheKey) (*Entry[ObjectData], error)

	// Cache stores an entry in the cache with the specified input key.
	Cache(key CacheKey, data io.Reader, expires time.Time, objectData ObjectData) (*Entry[ObjectData], error)

	// Delete removes an entry from the cache by its input key.
	Delete(key CacheKey) error

	// UpdateMetadata modifies the metadata of an entry in the cache.
	UpdateMetadata(key CacheKey, modifier func(*EntryMetadata[ObjectData])) error
}
