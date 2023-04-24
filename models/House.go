package models

import "gorm.io/gorm"

type House struct {
	gorm.Model

	Title        string `gorm:"not null" json:"title"`
	Price        string `gorm:"not null" json:"price"`
	Location     string `gorm:"not null" json:"location"`
	Description  string `gorm:"not null" json:"description"`
	PhotoURL     string `json:"photo_url"`
	Status       bool   `gorm:"default:true" json:"status"`
	AgentID      uint   `json:"agent_id"`
	HouseDetails HouseDetails
	HouseGallery HouseGallery
}
