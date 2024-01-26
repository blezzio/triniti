package usecases

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/blezzio/triniti/services/dtos"
	"github.com/blezzio/triniti/services/interfaces"
	"github.com/blezzio/triniti/utils"
)

type URL struct {
	repo   interfaces.URLRepo
	hash   interfaces.Hash
	logger interfaces.Logger
	cache  interfaces.Cache[string, string]
}

func NewURL(
	repo interfaces.URLRepo, hash interfaces.Hash, opts ...URLOpt,
) *URL {
	uc := &URL{repo: repo, hash: hash, logger: slog.Default()}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

func (uc *URL) GetHash(ctx context.Context, fullURL string) (string, error) {
	hash, err := uc.createNewHash(ctx, fullURL)
	if err := uc.setCache(ctx, hash, fullURL); err != nil {
		uc.logger.WarnContext(ctx, "failed to set cache", "key", hash, "value", fullURL, "error", err)
	}
	return hash, utils.Trace(err, "failed to get hash for url %s", fullURL)
}

func (uc *URL) createNewHash(ctx context.Context, fullURL string) (string, error) {
	getter := uc.hash.Hash(fullURL)
	var hash string
	for !getter.Done() {
		hash = getter.Get()
		existedURL, err := uc.GetFullURL(ctx, hash)
		if err != nil {
			break
		}
		if fullURL == existedURL {
			return hash, nil
		}
	}
	err := uc.repo.Create(ctx, &dtos.CreateHash{
		Hash:    hash,
		FullURL: fullURL,
	})

	return hash, utils.Trace(err, "failed to create hash for url %s", fullURL)
}

func (uc *URL) getCache(ctx context.Context, key string) (string, error) {
	if uc.cache == nil {
		return "", fmt.Errorf("no cache existed")
	}
	val, err := uc.cache.GetContext(ctx, key)
	trace := utils.Trace(err, "failed to get cache with key %q", key)
	return val, trace
}

func (uc *URL) setCache(ctx context.Context, key, value string) error {
	if uc.cache == nil {
		return fmt.Errorf("no cache existed")
	}
	err := uc.cache.SetContext(ctx, key, value)
	trace := utils.Trace(err, "failed to set cache with key=%s and value=%s", key, value)
	return trace
}

func (uc *URL) GetFullURL(ctx context.Context, hash string) (string, error) {
	fullURL, cacheErr := uc.getCache(ctx, hash)
	if cacheErr == nil {
		return fullURL, nil
	} else {
		uc.logger.WarnContext(ctx, "failed to get cache", "key", hash, "error", cacheErr)
	}

	fullURL, err := uc.repo.GetFullURL(ctx, hash)
	if err != nil {
		return "", utils.Trace(err, "failed to get full url for hash %s", hash)
	}

	if cacheErr == nil {
		if err := uc.setCache(ctx, hash, fullURL); err != nil {
			uc.logger.WarnContext(ctx, "failed to set cache", "key", hash, "value", fullURL, "error", err)
		}
	}

	return fullURL, nil
}
