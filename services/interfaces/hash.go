package interfaces

import "github.com/blezzio/triniti/services/dtos"

type Hash interface {
	Hash(string) *dtos.Hash
}
