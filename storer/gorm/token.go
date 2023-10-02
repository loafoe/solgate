package gorm

import (
	"github.com/loafoe/solgate/types"
	"gorm.io/gorm"
)

type TokenStorer struct {
	DB *gorm.DB
}

func (t TokenStorer) Create(token types.Token) (*types.Token, error) {
	tx := t.DB.Create(token)
	if tx.Error != nil {
		return nil, tx.Error
	}
	createdToken, err := t.FindByID(token.ID)
	return createdToken, err
}

func (t TokenStorer) Delete(id string) error {
	tx := t.DB.Delete(&types.Token{}, "id = ?", id)
	return tx.Error
}

func (t TokenStorer) FindByID(id string) (*types.Token, error) {
	var token types.Token
	tx := t.DB.First(&token, "id = ?", id)
	return &token, tx.Error
}

func (t TokenStorer) FindByToken(tkn string) (*types.Token, error) {
	var token types.Token
	tx := t.DB.First(&token, "value = ?", tkn)
	return &token, tx.Error
}
