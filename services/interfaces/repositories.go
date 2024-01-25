package interfaces

import (
	"context"

	"github.com/blezzio/tini/services/dtos"
)

type URLRepo interface {
	Create(ctx context.Context, param *dtos.CreateHash) error
	GetFullURL(ctx context.Context, hash string) (string, error)
	GetHash(ctx context.Context, fullURL string) (string, error)
	Delete(ctx context.Context, hashOrFullURL string) error
}
