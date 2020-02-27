//  tracer package add stack trace to the errors.
//  - Use trace Trace function to trace the error
//  - Use tracer Cause function to get he base error
//  - Use unwrap function from standard errors package to unwrap error.
//  - Use Is function from standard errors package to check error is expected error or no.
//
package tracer

// stack represents a stack of program counters.
type (
	tracedError struct {
		error
		*stack
	}

	// traceErr is the error struct that contain trace of error.
	StackTracer interface {
		StackTrace() StackTrace
	}
)

// Trace function check if error contains trace, so
// return it, otherwise add stacktrace to the error.
func Trace(err error) error {
	if err == nil {
		return nil
	}
	return &tracedError{
		err,
		callers(),
	}
}

func (e *tracedError) Cause() error { return e.error }

// Unwrap provides compatibility for Go 1.13 error chains.
func (e *tracedError) Unwrap() error { return e.error }

// Cause function return the base error that cause other errors.
func Cause(err error) error {
	if e, ok := err.(tracedError); err != nil && ok {
		return e.Cause()
	}

	return err
}

func MoveStack(from error, to error) error {
	tErr, ok := from.(tracedError)

	if from == nil || to == nil || !ok {
		return Trace(to)
	}

	return &tracedError{
		error: to,
		stack: tErr.stack,
	}
}
