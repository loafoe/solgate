package gorm

import (
	"github.com/loafoe/solgate/storer"
	"github.com/loafoe/solgate/types"
	"gorm.io/gorm"
)

func New(db *gorm.DB) storer.Solgate {
	_ = db.AutoMigrate(&types.Token{})

	return storer.Solgate{
		Token: &TokenStorer{DB: db},
	}
}
