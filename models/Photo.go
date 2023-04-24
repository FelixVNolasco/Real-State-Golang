package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model

	HouseGalleryID uint   `json:"house_gallery_id"`
	PhotoUrl       string `json:"url"`
}
