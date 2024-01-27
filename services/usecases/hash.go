package usecases

import (
	"encoding/base64"
	"hash"

	"github.com/blezzio/triniti/services/dtos"
)

type Hasher struct {
	h hash.Hash
}

func NewHasher(hasher hash.Hash) *Hasher {
	return &Hasher{h: hasher}
}

func (uc *Hasher) Hash(val string) *dtos.HashGetter {
	hash := base64.URLEncoding.EncodeToString(uc.h.Sum([]byte(val)))
	return dtos.NewHashGetter(hash)
}
