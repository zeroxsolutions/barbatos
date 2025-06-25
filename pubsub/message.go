package pubsub

// Message defines an interface for a publish-subscribe messaging system.
// It provides methods to retrieve the topic of the message and its associated data.
type Message interface {
	// Topic returns the topic or subject of the message.
	Topic() string

	// Data returns the payload of the message as a slice of bytes.
	Data() []byte
}
