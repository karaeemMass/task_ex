package repository

import (
	_ "task_ex/internal/model"

	"gorm.io/gorm"
)

type PriceGoldRepository struct {
	db *gorm.DB
}

func NewPriceGoldRepository(db *gorm.DB) *PriceGoldRepository {
	return &PriceGoldRepository{db: db}
}

