package models

import (
	"time"
)

type Post struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Topic     string `gorm:"not null; default:null"`
	Title     string `gorm:"not null; unique; default:null"`
	Body      string `gorm:"not null; default:null"`
	Image     []byte `gorm:"not null; default:null"`
	Abstract  string `gorm:"not null; default:null"`
}
