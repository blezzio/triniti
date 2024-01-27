package usecases

import (
	"hash"

	"github.com/blezzio/triniti/services/interfaces"
)

type HashOpt func(hasher *Hasher)

func WithHash(h hash.Hash) HashOpt {
	return func(hasher *Hasher) {
		hasher.h = h
	}
}

func WithEncoding(e interfaces.Encoding) HashOpt {
	return func(hasher *Hasher) {
		hasher.e = e
	}
}
