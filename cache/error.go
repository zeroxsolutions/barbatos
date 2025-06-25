package cache

import "errors"

// ErrCacheNil represents the error returned when a nil value is encountered
// during an operation in the cache package. This error is used to indicate that
// a requested key exists in the cache but its associated value is nil, which
// may imply that the key is present but uninitialized or cleared.
var ErrCacheNil = errors.New("cache: nil")
