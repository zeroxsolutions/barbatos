// Package pubsub provides a publish-subscribe system where publishers can send
// messages to topics, and subscribers can receive them. This package defines
// the Subscriber interface, which represents the behavior of a subscriber in
// the system, including subscribing, unsubscribing, receiving messages, checking
// the connection status, and closing the subscriber.
package pubsub

import "context"

// Subscriber defines an interface for a subscriber in a publish-subscribe system.
// It provides methods to subscribe to and unsubscribe from topics, receive messages,
// check the connection status, and close the subscriber.
//
// A Subscriber is expected to have the following behaviors:
// - It can subscribe to one or more topics.
// - It can unsubscribe from topics.
// - It provides a receiver channel to fetch messages.
// - It can check whether the subscriber is connected to the pub-sub system.
// - It can be closed to release any associated resources.
type Subscriber interface {
	// Subscribe subscribes the subscriber to one or more topics.
	// It accepts a context to handle timeouts or cancellations.
	// Returns an error if the subscription fails (e.g., invalid topic or connection issue).
	//
	// Example:
	//     err := subscriber.Subscribe(ctx, "topic1", "topic2")
	Subscribe(ctx context.Context, topics ...string) error

	// Unsubscribe removes the subscriber's subscription from one or more topics.
	// It accepts a context and returns an error if the operation fails (e.g., topic not found).
	//
	// Example:
	//     err := subscriber.Unsubscribe(ctx, "topic1")
	Unsubscribe(ctx context.Context, topics ...string) error

	// Receiver returns a channel from which messages can be received.
	// It accepts a context to allow cancellation or timeout during the message retrieval.
	// The returned channel is receive-only and carries Message objects.
	// If the operation fails, it returns an error.
	//
	// Example:
	//     messages, err := subscriber.Receiver(ctx)
	//     for msg := range messages {
	//         // handle msg
	//     }
	Receiver(ctx context.Context) (<-chan Message, error)

	// IsConnected checks if the subscriber is currently connected to the pub-sub system.
	// It accepts a context and returns true if the subscriber is connected, otherwise false.
	// This method helps to monitor the connection status of the subscriber.
	//
	// Example:
	//     if subscriber.IsConnected(ctx) {
	//         // handle connected state
	//     }
	IsConnected(ctx context.Context) bool

	// Close closes the subscriber and releases any resources associated with it.
	// Returns an error if the operation fails (e.g., already closed).
	//
	// Example:
	//     err := subscriber.Close()
	Close() error
}
