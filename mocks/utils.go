package mocks

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type callLog struct {
	callMap   map[string][]any
	callCount map[string]int
}

func (c callLog) insertCallLog(params ...any) {
	fullFuncName := c.getCallerName()
	parts := strings.Split(fullFuncName, ".")
	funcName := parts[len(parts)-1]
	c.callMap[funcName] = params
	c.callCount[funcName]++
}

func (c callLog) Called(f any) (int, []any) {
	fullFuncName := c.getFuncName(f)
	parts := strings.Split(fullFuncName, ".")

	infFuncName := parts[len(parts)-1]
	funcName := strings.Replace(infFuncName, "-fm", "", 1)

	return c.callCount[funcName], c.callMap[funcName]
}

func (c callLog) getCallerName() string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		return fmt.Sprintf("%s.(%d)", file, line)
	}
	return f.Name()
}

func (c callLog) getFuncName(f any) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
