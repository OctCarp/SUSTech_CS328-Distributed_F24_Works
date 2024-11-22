package models

import "time"

type Order struct {
	ID         int32     `gorm:"primaryKey"`
	UserID     int32     `gorm:"not null"`
	ProductID  int32     `gorm:"not null"`
	Quantity   int32     `gorm:"check:quantity > 0 AND quantity <= 3"`
	TotalPrice float64   `gorm:"type:decimal(10,2);not null"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	User       User      `gorm:"foreignKey:UserID"`
	Product    Product   `gorm:"foreignKey:ProductID"`
}
