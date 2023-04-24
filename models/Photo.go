package models

import "gorm.io/gorm"

//TODO: FIX ERROR: relation "photos" does not exist (SQLSTATE 42P01)
type Photo struct {
	gorm.Model

	HouseGalleryID uint `json:"house_gallery_id"`
	url            string
}
