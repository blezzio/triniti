package interfaces

import "context"

type Cache[K comparable, V any] interface {
	Set(key K, value V) error
	Get(key K) (V, error)
	Len() int
	SetContext(ctx context.Context, key K, value V) error
	GetContext(ctx context.Context, key K) (V, error)
	LenContext(ctx context.Context) int
}
