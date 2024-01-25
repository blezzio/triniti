// go:build mocks
package mocks

import (
	"context"
	"fmt"
)

type TestLogger struct {
	callLog
	logs []string
}

func NewTestLogger() *TestLogger {
	return &TestLogger{
		logs: []string{},
		callLog: callLog{
			callMap:   map[string][]any{},
			callCount: map[string]int{},
		},
	}
}

func (l *TestLogger) Info(msg string, args ...any) {
	l.insertCallLog(msg, args)
	l.logs = append(l.logs, fmt.Sprintf(msg, args...))
}

func (l *TestLogger) Warn(msg string, args ...any) {
	l.insertCallLog(msg, args)
	l.Info(msg, args...)
}

func (l *TestLogger) Error(msg string, args ...any) {
	l.insertCallLog(msg, args)
	l.Info(msg, args...)
}

func (l *TestLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.insertCallLog(msg, args)
	l.Info(msg, args...)
}

func (l *TestLogger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.insertCallLog(msg, args)
	l.Info(msg, args...)
}

func (l *TestLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.insertCallLog(msg, args)
	l.Info(msg, args...)
}
