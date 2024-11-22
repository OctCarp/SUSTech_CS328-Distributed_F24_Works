package models

import "time"

type Product struct {
	ID          int32     `gorm:"primaryKey"`
	Name        string    `gorm:"size:100;not null"`
	Description string    `gorm:"type:text"`
	Category    string    `gorm:"size:50"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	Slogan      string    `gorm:"size:255"`
	Stock       int32     `gorm:"not null;default:500"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
