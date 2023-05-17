package models

import "time"

type Subscriber struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string `gorm:"unique"`
	Name      string `gorm:"not null"`
}
