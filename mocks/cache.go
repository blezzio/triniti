// go:build mocks
package mocks

import (
	"context"
	"fmt"
)

type TestCache struct {
	callLog
	item    map[string]string
	withErr error
}

func NewTestCache(item map[string]string, withErr error) *TestCache {
	return &TestCache{
		item:    item,
		withErr: withErr,
		callLog: callLog{
			callMap:   map[string][]any{},
			callCount: map[string]int{},
		},
	}
}

func (c *TestCache) Set(key string, value string) error {
	c.insertCallLog(key, value)
	if c.withErr != nil {
		return c.withErr
	}
	c.item[key] = value
	return nil
}

func (c *TestCache) Get(key string) (string, error) {
	c.insertCallLog(key)
	if c.withErr != nil {
		return "", c.withErr
	}
	val, ok := c.item[key]
	if !ok {
		return "", fmt.Errorf("not found")
	}
	return val, nil

}

func (c *TestCache) Len() int {
	c.insertCallLog()
	return len(c.item)
}

func (c *TestCache) SetContext(ctx context.Context, key string, value string) error {
	c.insertCallLog(ctx, key, value)
	return c.Set(key, value)
}

func (c *TestCache) GetContext(ctx context.Context, key string) (string, error) {
	c.insertCallLog(ctx, key)
	return c.Get(key)
}

func (c *TestCache) LenContext(ctx context.Context) int {
	c.insertCallLog(ctx)
	return c.Len()
}
