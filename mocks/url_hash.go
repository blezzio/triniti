// go:build mocks
package mocks

import "github.com/blezzio/triniti/services/dtos"

type URLTestHash struct {
	callLog
}

func NewURLTestHash() *URLTestHash {
	return &URLTestHash{
		callLog: callLog{
			callMap:   map[string][]any{},
			callCount: map[string]int{},
		},
	}
}

func (h *URLTestHash) Hash(val string) *dtos.Hash {
	h.insertCallLog(val)
	return dtos.NewHash("abcdef")
}
