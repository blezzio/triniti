package interfaces

import "context"

type URLUseCase interface {
	GetHash(ctx context.Context, fullURL string) (string, error)
	GetFullURL(ctx context.Context, hash string) (string, error)
}
