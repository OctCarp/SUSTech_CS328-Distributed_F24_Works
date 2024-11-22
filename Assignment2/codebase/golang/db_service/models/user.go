package models

import "time"

type User struct {
	ID           int32  `gorm:"primaryKey"`
	SID          string `gorm:"column:sid;size:15;uniqueIndex;not null"`
	Username     string `gorm:"size:50;uniqueIndex;not null"`
	Email        string `gorm:"size:255;uniqueIndex"`
	PasswordHash string `gorm:"size:255;not null"`
	CreatedAt    time.Time
}
