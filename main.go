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

	// Agents Routes

	r.HandleFunc("/agents", routes.GetAgentsHandler).Methods("GET")
	r.HandleFunc("/agents/{id}", routes.GetAgentHandler).Methods("GET")
	r.HandleFunc("/agents", routes.PostAgentsHandler).Methods("POST")
	r.HandleFunc("/agents/{id}", routes.UpdateAgentHandler).Methods("PUT")
	r.HandleFunc("/agents/{id}", routes.DeleteAgentHandler).Methods("DELETE")

	// Houses Routes

	r.HandleFunc("/houses", routes.GetHousesHandler).Methods("GET")
	r.HandleFunc("/houses/{id}", routes.GetHouseHandler).Methods("GET")
	r.HandleFunc("/houses", routes.PostHouseHandler).Methods("POST")
	r.HandleFunc("/houses/{id}", routes.UpdateHouseHandler).Methods("POST")
	r.HandleFunc("/houses/{id}", routes.DeleteHousesHandler).Methods("DELETE")

	// House Details Routes

	r.HandleFunc("/house/details/{id}", routes.GetHouseDetailsHandler).Methods("GET")
	r.HandleFunc("/house/details", routes.PostHouseDetailsHandler).Methods("POST")
	r.HandleFunc("/house/details/{id}", routes.UpdateHouseDetailsHandler).Methods("PUT")
	r.HandleFunc("/house/details/{id}", routes.DeleteHouseDetailsHandler).Methods("DELETE")

	// House Gallery Routes

	r.HandleFunc("/gallery/{id}", routes.GetHouseGalleryHandler).Methods("GET")
	r.HandleFunc("/gallery", routes.PostHouseGalleryHandler).Methods("POST")

	// House Gallery Routes

	r.HandleFunc("/photo/{id}", routes.GetPhotoHandler).Methods("GET")
	r.HandleFunc("/photo", routes.PostPhotoHandler).Methods("POST")

	http.ListenAndServe(":3000", r)
}
