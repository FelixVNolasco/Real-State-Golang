package main

import (
	"log"
	"net/http"

	"github.com/FelixVNolasco/real-state-golang/db"
	"github.com/FelixVNolasco/real-state-golang/models"
	"github.com/FelixVNolasco/real-state-golang/routes"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Agent{})
	db.DB.AutoMigrate(models.House{})
	db.DB.AutoMigrate(models.HouseDetails{})
	db.DB.AutoMigrate(models.HouseGallery{})
	db.DB.AutoMigrate(models.Photo{})

	// Create an API handler which serves data from PlanetScale.
	handler := routes.NewHandler(db.DBConnection())

	const addr = ":8080"

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
