package storer

import (
	"github.com/loafoe/solgate/types"
)

type Token interface {
	Create(token types.Token) (*types.Token, error)
	Delete(id string) error
	FindByID(id string) (*types.Token, error)
}
