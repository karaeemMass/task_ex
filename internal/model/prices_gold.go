package model


type PricesGold struct {
	ID               uint `gorm:"primaryKey"`
	PriceMain        float64
	PriceBefore      float64
	Cunt             string
	Date             int64
	BaseCurrency     string
	WeightUnit       string
	WeightName       string
	Open             float64
	High             float64
	Low              float64
	Prev             float64
	Change           float64
	ChangePercentage float32
	Price24k         float64
	Price22k         float64
	Price21k         float64
	Price18k         float64
}
