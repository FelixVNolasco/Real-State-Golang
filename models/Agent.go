package models

import "gorm.io/gorm"

type Agent struct {
	gorm.Model

	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	PhoneNumber string `gorm:"not null"`
	Email       string `gorm:"not null"`
	Houses      []House
}
