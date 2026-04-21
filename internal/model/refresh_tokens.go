package model

import "gorm.io/gorm"

type RefreshToken struct {
	gorm.Model

	UserID    uint   `gorm:"not null"`
	TokenHash string `gorm:"size:255;not null"`
	ExpiresAt int64  `gorm:"not null"`
	CreatedAt int64  `gorm:"not null"`
}
