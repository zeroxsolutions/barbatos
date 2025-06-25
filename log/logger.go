package log

// Logger defines an interface for a generic logging system, providing methods for logging
// messages at various severity levels: Debug, Info, Warn, Error, Panic, and Fatal.
// Each level supports multiple logging styles, including unformatted logs, formatted logs
// (with templates), and structured logs (with key-value pairs).
type Logger interface {
	// Debug logs debug messages without any specific formatting.
	// This is typically used for detailed troubleshooting information during development.
	Debug(args ...interface{})

	// Debugf logs formatted debug messages using a format template and additional arguments.
	// Allows for more controlled logging of dynamic debug information.
	Debugf(template string, args ...interface{})

	// Debugw logs debug messages with additional key-value pairs for structured logging.
	// This is useful for logging context information alongside the debug message.
	Debugw(msg string, keysValues ...interface{})

	// Info logs informational messages without any specific formatting.
	// Use this for general runtime information about the system's normal behavior.
	Info(args ...interface{})

	// Infof logs formatted informational messages using a format template and additional arguments.
	// Similar to Info, but allows for more detailed, dynamic information.
	Infof(template string, args ...interface{})

	// Infow logs informational messages with additional key-value pairs for structured logging.
	// Structured logging allows for easier querying and analysis of logs.
	Infow(msg string, keysValues ...interface{})

	// Warn logs warning messages without any specific formatting.
	// Used to highlight situations that might lead to errors but haven't yet.
	Warn(args ...interface{})

	// Warnf logs formatted warning messages using a format template and additional arguments.
	// Allows logging of more dynamic warning details.
	Warnf(template string, args ...interface{})

	// Warnw logs warning messages with additional key-value pairs for structured logging.
	// This is useful when warnings need to be accompanied by additional context.
	Warnw(msg string, keysValues ...interface{})

	// Error logs error messages without any specific formatting.
	// This is used to record runtime errors that are handled but still important to log.
	Error(args ...interface{})

	// Errorf logs formatted error messages using a format template and additional arguments.
	// Useful for logging errors with dynamic data.
	Errorf(template string, args ...interface{})

	// Errorw logs error messages with additional key-value pairs for structured logging.
	// This allows capturing error context in a structured manner.
	Errorw(msg string, keysValues ...interface{})

	// Panic logs panic messages without any specific formatting and triggers a panic.
	// A panic will stop the normal execution of the program and start unwinding the stack.
	Panic(args ...interface{})

	// Panicf logs formatted panic messages using a format template and additional arguments,
	// then triggers a panic.
	Panicf(template string, args ...interface{})

	// Panicw logs panic messages with additional key-value pairs for structured logging
	// and then triggers a panic.
	Panicw(msg string, keysValues ...interface{})

	// Fatal logs fatal messages without any specific formatting and terminates the program.
	// A fatal log indicates a severe issue that requires immediate program termination.
	Fatal(args ...interface{})

	// Fatalf logs formatted fatal messages using a format template and additional arguments,
	// then terminates the program.
	Fatalf(template string, args ...interface{})

	// Fatalw logs fatal messages with additional key-value pairs for structured logging
	// and then terminates the program.
	Fatalw(msg string, keysValues ...interface{})
}
