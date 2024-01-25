package usecases

import "github.com/blezzio/tini/services/interfaces"

type URLOpt func(*URL)

func WithCache(cache interfaces.Cache[string, string]) URLOpt {
	return func(uc *URL) {
		uc.cache = cache
	}
}

func WithLogger(logger interfaces.Logger) URLOpt {
	return func(uc *URL) {
		uc.logger = logger
	}
}
