package routes

import (
	"net/http"

	"github.com/FelixVNolasco/real-state-golang/db"
	"github.com/FelixVNolasco/real-state-golang/models"
)

func seedDatabase(w http.ResponseWriter, r *http.Request) {

	// db.DB.AutoMigrate(models.Agent{})
	// db.DB.AutoMigrate(models.House{})
	// db.DB.AutoMigrate(models.HouseDetails{})
	// db.DB.AutoMigrate(models.HouseGallery{})
	// db.DB.AutoMigrate(models.Photo{})

	// Perform initial schema migrations.
	if err := db.DB.AutoMigrate(models.Agent{}); err != nil {
		http.Error(w, "failed to migrate Agent table", http.StatusInternalServerError)
		return
	}
	if err := db.DB.AutoMigrate(models.House{}); err != nil {
		http.Error(w, "failed to migrate House table", http.StatusInternalServerError)
		return
	}
	if err := db.DB.AutoMigrate(models.HouseDetails{}); err != nil {
		http.Error(w, "failed to migrate HouseDetails table", http.StatusInternalServerError)
		return
	}
	if err := db.DB.AutoMigrate(models.HouseGallery{}); err != nil {
		http.Error(w, "failed to migrate HouseGallery table", http.StatusInternalServerError)
		return
	}
	if err := db.DB.AutoMigrate(models.Photo{}); err != nil {
		http.Error(w, "failed to migrate Photo table", http.StatusInternalServerError)
		return
	}
}
