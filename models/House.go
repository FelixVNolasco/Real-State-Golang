package models

import "gorm.io/gorm"

type House struct {
	gorm.Model

	Title       string `gorm:"not null"`
	Price       string `gorm:"not null"`
	Location    string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status      bool   `gorm:"default:false"`
	AgentID     uint
}
