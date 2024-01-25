package utils

import (
	"fmt"
	"runtime"
)

type terr struct {
	cause error
	msg   string
}

func (err terr) Error() string {
	return fmt.Sprintf("%s\ncaused by %v\n", err.msg, err.cause.Error())
}

func Trace(cause error, msg string, args ...any) error {
	if cause == nil {
		return nil
	}

	desc := fmt.Sprintf(msg, args...)
	err := terr{
		cause: cause,
		msg:   desc,
	}

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return err
	}
	err.msg = fmt.Sprintf("error at line %d in file %s:\n\t%s", line, file, desc)

	f := runtime.FuncForPC(pc)
	if f == nil {
		return err
	}
	err.msg = fmt.Sprintf("error at line %d in file %s(%s):\n\t%s", line, file, f.Name(), desc)

	return err
}
