package tracer

// Trace right now just simply returns the error with its stack.
// But later we may change its behaviour to be configurable, for
// example do nothing in production, or just wrap with
// file_name, line number of where the error occurred instead of
// full stack.
func Trace(err error) error { return WithStack(err) }

// TraceFrom moves the trace of the error or if the old error doesn't have any trace,
// it'll generate new trace.
func TraceFrom(from error, to error) error {
	return StackFrom(from, to)
}
