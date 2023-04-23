package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixVNolasco/real-state-golang/db"
	"github.com/FelixVNolasco/real-state-golang/models"
	"github.com/gorilla/mux"
)

func GetHousesHandler(w http.ResponseWriter, r *http.Request) {
	var houses []models.House
	db.DB.Find(&houses)
	json.NewEncoder(w).Encode(houses)
}

func PostHouseHandler(w http.ResponseWriter, r *http.Request) {
	var house models.House
	json.NewDecoder(r.Body).Decode(&house)

	createdHouse := db.DB.Create(&house)
	err := createdHouse.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&house)
}

func GetHouseHandler(w http.ResponseWriter, r *http.Request) {

	var house models.House
	params := mux.Vars(r)

	db.DB.First(&house, params["id"])

	if house.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("House Not Found"))
		return
	}

	json.NewEncoder(w).Encode(&house)

}

func DeleteHousesHandler(w http.ResponseWriter, r *http.Request) {

	var house models.House
	params := mux.Vars(r)

	db.DB.First(&house, params["id"])

	if house.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("House Not Found"))
		return
	}

	db.DB.Delete(&house)
	w.WriteHeader(http.StatusNoContent)

}
