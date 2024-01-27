package utils

import (
	"errors"
	"fmt"
)

type TraceError struct {
	cause error
	msg   error
}

func (err TraceError) Is(target error) bool {
	if errors.Is(err.msg, target) {
		return true
	}
	if err.cause == nil {
		return false
	}

	cause, isTraceErr := err.cause.(TraceError)
	if !isTraceErr {
		return errors.Is(err.cause, target)
	}
	return cause.Is(target)
}

func (err TraceError) Error() string {
	return fmt.Sprintf("%s\ncaused by %s\n", err.msg, err.cause)
}
