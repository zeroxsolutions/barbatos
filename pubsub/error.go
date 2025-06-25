// Package pubsub provides utilities related to publish-subscribe messaging patterns.
// This package defines common errors encountered in pub-sub operations.
package pubsub

import "errors"

// ErrConnectFailed is returned when a connection attempt to the pub-sub system fails.
// This error indicates that the client could not establish a connection to the server
// or broker, potentially due to network issues, incorrect credentials, or misconfigured endpoints.
var ErrConnectFailed = errors.New("pubsub: connect is failed")
