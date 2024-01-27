package utils

import (
	"fmt"
	"testing"
)

func TestTraceError(t *testing.T) {
	rootErr := fmt.Errorf("root error")
	err := Trace(rootErr, "first failure")
	err = Trace(err, "trace log")
	err = Trace(err, "trace log")
	terr, isTraceErr := err.(TraceError)
	if !isTraceErr {
		t.Error("err is not of type TraceError")
	}
	if !terr.Is(rootErr) {
		t.Errorf("err is not %v", rootErr)
	}

	firstErr := fmt.Errorf("first error")
	secondErr := fmt.Errorf("second error")
	err = Trace(firstErr, "failed, additionally %w", secondErr)
	terr, isTraceErr = err.(TraceError)
	if !isTraceErr {
		t.Error("err is not of type TraceError")
	}
	if !terr.Is(firstErr) {
		t.Errorf("err is not %v", firstErr)
	}
	if !terr.Is(secondErr) {
		t.Errorf("err is not %v", secondErr)
	}
}
