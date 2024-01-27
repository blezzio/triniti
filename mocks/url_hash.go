// go:build mocks
package mocks

import "github.com/blezzio/triniti/services/dtos"

type URLTestHash struct {
	callLog
	r string
}

func NewURLTestHash(result string) *URLTestHash {
	return &URLTestHash{
		callLog: callLog{
			callMap:   map[string][]any{},
			callCount: map[string]int{},
		},
		r: result,
	}
}

func (h *URLTestHash) Hash(val string) *dtos.HashGetter {
	h.insertCallLog(val)
	return dtos.NewHashGetter(h.r)
}
