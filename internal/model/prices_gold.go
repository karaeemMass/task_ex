package model

//import "gorm.io/gorm"

type price struct {
	ID                uint `gorm:"primaryKey"`
	PriceMain         float64
	PriceBefore       float64
	Cunt              string
	Date              int64
	BaseCurrency      string
	WeightUnit        string
	weightName        string
	open              float64
	high              float64
	low               float64
	prev              float64
	change            float64
	change_percentage float32
	price_24k         float64
	price_22k         float64
	price_21k         float64
	price_20k         float64
}
