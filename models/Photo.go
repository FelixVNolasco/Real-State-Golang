package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model

	url            string
	HouseGalleryID uint
}
