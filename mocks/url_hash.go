// go:build mocks
package mocks

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

func (h *URLTestHash) Hash(val string) string {
	h.insertCallLog(val)
	return "abcdef"
}
