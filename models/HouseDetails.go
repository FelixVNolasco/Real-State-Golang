package models

import "gorm.io/gorm"

type HouseDetails struct {
	gorm.Model

	HouseID    uint   `json:"house_id"`
	Rooms      int    `json:"rooms"`
	Bathrooms  string `json:"bathrooms"`
	ParkingLot string `json:"parking_lot"`
	SQ_MT      string `json:"sq_mt"`
}
