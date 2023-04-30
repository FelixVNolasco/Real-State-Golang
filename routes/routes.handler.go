package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// A Handler is an HTTP API server handler.
type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) http.Handler {

	r := mux.NewRouter()

	// AutomMigrate
	r.HandleFunc("/automigrate", seedDatabase).Methods("GET")

	// Agents Routes
	r.HandleFunc("/agents", GetAgentsHandler).Methods("GET")
	r.HandleFunc("/agents/{id}", GetAgentHandler).Methods("GET")
	r.HandleFunc("/agents", PostAgentsHandler).Methods("POST")
	r.HandleFunc("/agents/{id}", UpdateAgentHandler).Methods("PUT")
	r.HandleFunc("/agents/{id}", DeleteAgentHandler).Methods("DELETE")

	// Houses Routes
	r.HandleFunc("/houses", GetHousesHandler).Methods("GET")
	r.HandleFunc("/houses/{id}", GetHouseHandler).Methods("GET")
	r.HandleFunc("/houses", PostHouseHandler).Methods("POST")
	r.HandleFunc("/houses/{id}", UpdateHouseHandler).Methods("PUT")
	r.HandleFunc("/houses/{id}", DeleteHousesHandler).Methods("DELETE")

	// House Details Routes
	r.HandleFunc("/house/details/{id}", GetHouseDetailsHandler).Methods("GET")
	r.HandleFunc("/house/details", PostHouseDetailsHandler).Methods("POST")
	r.HandleFunc("/house/details/{id}", UpdateHouseDetailsHandler).Methods("PUT")
	r.HandleFunc("/house/details/{id}", DeleteHouseDetailsHandler).Methods("DELETE")

	// House Gallery Routes
	r.HandleFunc("/gallery/{id}", GetHouseGalleryHandler).Methods("GET")
	r.HandleFunc("/gallery", PostHouseGalleryHandler).Methods("POST")

	// Photo Routes
	r.HandleFunc("/photo/{id}", GetPhotoHandler).Methods("GET")
	r.HandleFunc("/photo", PostPhotoHandler).Methods("POST")

	return r
}
