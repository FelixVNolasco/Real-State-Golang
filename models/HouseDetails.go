package models

import "gorm.io/gorm"

type HouseDetails struct {
	gorm.Model

	HouseID    uint
	Rooms      int
	Bathrooms  string
	ParkingLot string
	SQ_MT      string
}
