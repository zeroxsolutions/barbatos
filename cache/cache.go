package cache

import (
	"context"
	"time"
)

// Cache represents the interface for cache operations.
// It defines a set of methods for interacting with a cache system,
// such as checking connectivity, retrieving, storing, and deleting values.
type Cache interface {

	// IsConnected checks whether the connection to the cache system is established.
	// It returns true if the connection is active, and false otherwise.
	IsConnected(ctx context.Context) bool

	// Keys returns a list of keys in the cache system that match the given pattern.
	// The pattern allows for wildcard searches (e.g., "*"), and the function
	// returns the matching keys and any error encountered during the operation.
	Keys(ctx context.Context, pattern string) ([]string, error)

	// Get retrieves the value associated with the given key from the cache system.
	// It returns the value as a string and any error encountered during the operation.
	Get(ctx context.Context, key string) (string, error)

	// Set stores a value in the cache system with the specified key.
	// It accepts the key and the value to be stored. If the operation fails,
	// it returns an error.
	Set(ctx context.Context, key string, value interface{}) error

	// SetWithExpiration stores a value in the cache system with the specified key
	// and sets an expiration time for the key-value pair. After the expiration
	// duration, the value will be automatically removed from the cache.
	SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error

	// Del deletes the specified keys from the cache system. If the operation fails,
	// it returns an error. It accepts multiple keys as variadic arguments.
	Del(ctx context.Context, keys ...string) error

	// DelWithPattern deletes all keys that match the given pattern from the cache system.
	// It returns an error if the operation fails. This is useful for bulk deletion
	// of keys that share a common prefix or pattern.
	DelWithPattern(ctx context.Context, pattern string) error

	// Close closes the connection to the cache system. This should be called
	// when the cache client is no longer needed to release any resources held by it.
	Close() error
}
