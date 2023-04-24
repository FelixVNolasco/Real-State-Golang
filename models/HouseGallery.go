package models

import "gorm.io/gorm"

type HouseGallery struct {
	gorm.Model

	HouseID uint    `json:"house_id"`
	Photos  []Photo `json:"photos"`
}
