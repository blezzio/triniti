package dtos

import "github.com/blezzio/triniti/services/consts"

type HashGetter struct {
	val     string
	maxlen  int
	minlen  int
	currlen int
}

func NewHashGetter(val string) *HashGetter {
	maxlen := len(val)
	if consts.MaxHashLen < maxlen {
		maxlen = consts.MaxHashLen
	}
	return &HashGetter{
		val:     val,
		maxlen:  maxlen,
		minlen:  consts.MinHashLen,
		currlen: consts.MinHashLen,
	}
}

func (dto *HashGetter) Get() string {
	currlen := dto.currlen
	if dto.maxlen < currlen {
		currlen = dto.maxlen
	}
	val := dto.val[:currlen]

	dto.currlen++
	return val
}

func (dto *HashGetter) Done() bool {
	return dto.currlen > dto.maxlen
}
