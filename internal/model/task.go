package model

type Task struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	Completed   bool
}
