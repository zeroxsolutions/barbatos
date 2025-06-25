package pubsub

import "context"

// Publisher defines an interface for a publish-subscribe system's publisher.
// It provides methods for publishing messages to a topic, checking the connection status,
// and closing the publisher.
type Publisher interface {
	// Publish sends the provided messages to the specified topic.
	// It accepts a context for handling timeouts or cancellations.
	// Returns an error if the operation fails.
	Publish(ctx context.Context, topic string, messages ...[]byte) error

	// IsConnected checks if the publisher is currently connected to the pub-sub system.
	// It accepts a context and returns true if connected, otherwise false.
	IsConnected(ctx context.Context) bool

	// Close closes the publisher and releases any resources.
	// Returns an error if the operation fails.
	Close() error
}
