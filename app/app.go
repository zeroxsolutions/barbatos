package app

// App defines an interface for managing the application's lifecycle.
// It provides methods for running and shutting down the application,
// allowing for a structured approach to application control.
type App interface {
	// Run starts the application, initializing all necessary components,
	// configurations, and resources required for the application to operate.
	// It returns an error if the application fails to start correctly.
	Run() error

	// Shutdown gracefully terminates the application, handling cleanup
	// of resources and any necessary shutdown procedures. It ensures that
	// the application stops in a controlled manner, returning an error
	// if the shutdown process encounters issues.
	Shutdown() error
}
