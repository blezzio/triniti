package usecases

import (
	"encoding/base64"
	"hash"
	"sync"

	"github.com/blezzio/triniti/services/dtos"
)

type Hasher struct {
	mux sync.Mutex
	h   hash.Hash
}

func NewHasher(hasher hash.Hash) *Hasher {
	return &Hasher{mux: sync.Mutex{}, h: hasher}
}

func (uc *Hasher) Hash(val string) *dtos.HashGetter {
	uc.mux.Lock()
	defer uc.mux.Unlock()
	defer uc.h.Reset()

	uc.h.Write([]byte(val))
	hash := base64.URLEncoding.EncodeToString(uc.h.Sum(nil))

	return dtos.NewHashGetter(hash)
}
