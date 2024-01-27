package utils

import (
	"fmt"
	"runtime"
)

func Trace(cause error, msg string, args ...any) error {
	if cause == nil {
		return nil
	}

	desc := fmt.Errorf(msg, args...)
	err := TraceError{
		cause: cause,
		msg:   desc,
	}

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return err
	}
	err.msg = fmt.Errorf("error at line %d in file %s:\n\t%w", line, file, desc)

	f := runtime.FuncForPC(pc)
	if f == nil {
		return err
	}
	err.msg = fmt.Errorf("error at line %d in file %s(%s):\n\t%w", line, file, f.Name(), desc)

	return err
}
