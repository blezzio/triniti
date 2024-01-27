package usecases

import (
	"crypto/sha1"
	"encoding/base64"
	"hash"
	"sync"

	"github.com/blezzio/triniti/services/dtos"
	"github.com/blezzio/triniti/services/interfaces"
)

type Hasher struct {
	mux sync.Mutex
	h   hash.Hash
	e   interfaces.Encoding
}

func NewHasher(opts ...HashOpt) *Hasher {
	def := &Hasher{
		mux: sync.Mutex{},
		h:   sha1.New(),
		e:   base64.URLEncoding,
	}

	for _, opt := range opts {
		opt(def)
	}

	return def
}

func (uc *Hasher) Hash(val string) *dtos.HashGetter {
	uc.mux.Lock()
	defer uc.mux.Unlock()
	defer uc.h.Reset()

	uc.h.Write([]byte(val))
	hash := uc.e.EncodeToString(uc.h.Sum(nil))

	return dtos.NewHashGetter(hash)
}
