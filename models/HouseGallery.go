package models

import "gorm.io/gorm"

type HouseGallery struct {
	gorm.Model

	HouseID uint
	Photos  []Photo
}
