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

	// Agents Routes

	r.HandleFunc("/agents", routes.GetAgentsHandler).Methods("GET")
	r.HandleFunc("/agents/{id}", routes.GetAgentHandler).Methods("GET")
	r.HandleFunc("/agents", routes.PostAgentsHandler).Methods("POST")
	r.HandleFunc("/agents/{id}", routes.DeleteAgentHandler).Methods("DELETE")

	// Houses Routes

	r.HandleFunc("/houses", routes.GetHousesHandler).Methods("GET")
	r.HandleFunc("/houses/{id}", routes.GetHouseHandler).Methods("GET")
	r.HandleFunc("/houses", routes.PostHouseHandler).Methods("POST")
	r.HandleFunc("/houses/{id}", routes.DeleteHousesHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
