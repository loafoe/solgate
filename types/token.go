package types

import (
	"time"
)

type Token struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Value     string    `json:"value"`
	Endpoint  string    `json:"endpoint" form:"endpoint" gorm:"index"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
