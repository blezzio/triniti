package mocks

import (
	"context"
	"fmt"

	"github.com/blezzio/triniti/services/dtos"
)

type URLTestRepo struct {
	callLog
	item      map[string]string
	withError error
}

func NewUrlTestRepo(
	item map[string]string, withError error,
) *URLTestRepo {
	return &URLTestRepo{
		item:      item,
		withError: withError,
		callLog: callLog{
			callMap:   map[string][]any{},
			callCount: map[string]int{},
		},
	}
}

func (r *URLTestRepo) Create(ctx context.Context, param *dtos.CreateHash) error {
	r.insertCallLog(ctx, param)
	return r.withError
}

func (r *URLTestRepo) GetFullURL(ctx context.Context, hash string) (string, error) {
	r.insertCallLog(ctx, hash)
	if r.withError != nil {
		return "", r.withError
	}
	fullURL, ok := r.item[hash]
	if !ok {
		return "", fmt.Errorf("cannot find url")
	}
	return fullURL, nil
}

func (r *URLTestRepo) GetHash(ctx context.Context, fullURL string) (string, error) {
	r.insertCallLog(ctx, fullURL)
	if r.withError != nil {
		return "", r.withError
	}
	hash, ok := r.item[fullURL]
	if !ok {
		return "", fmt.Errorf("cannot find hash")
	}
	return hash, nil
}

func (r *URLTestRepo) Delete(ctx context.Context, hashOrFullURL string) error {
	r.insertCallLog(ctx, hashOrFullURL)
	return r.withError
}
