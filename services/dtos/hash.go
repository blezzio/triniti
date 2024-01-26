package dtos

import "github.com/blezzio/triniti/services/consts"

type Hash struct {
	val     string
	maxlen  int
	minlen  int
	currlen int
}

func NewHash(val string) *Hash {
	maxlen := len(val)
	if consts.MaxHashLen < maxlen {
		maxlen = consts.MaxHashLen
	}
	return &Hash{
		val:     val,
		maxlen:  maxlen,
		minlen:  consts.MinHashLen,
		currlen: consts.MinHashLen,
	}
}

func (dto *Hash) Next() string {
	currlen := dto.currlen
	if dto.maxlen < currlen {
		currlen = dto.maxlen
	}
	val := dto.val[:currlen]

	dto.currlen++
	return val
}

func (dto *Hash) Done() bool {
	return dto.currlen > dto.maxlen
}
