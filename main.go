package main

import (
	"net/http"

	"github.com/FelixVNolasco/real-state-golang/db"
	"github.com/FelixVNolasco/real-state-golang/models"
	"github.com/FelixVNolasco/real-state-golang/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Agent{})
	db.DB.AutoMigrate(models.House{})
	db.DB.AutoMigrate(models.HouseDetails{})
	db.DB.AutoMigrate(models.HouseGallery{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
