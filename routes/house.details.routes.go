package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixVNolasco/real-state-golang/db"
	"github.com/FelixVNolasco/real-state-golang/models"
	"github.com/gorilla/mux"
)

func GetHouseDetailsHandler(w http.ResponseWriter, r *http.Request) {

	var HouseDetails models.HouseDetails
	params := mux.Vars(r)

	db.DB.First(&HouseDetails, params["id"])

	if HouseDetails.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("House Not Found"))
		return
	}

	json.NewEncoder(w).Encode(&HouseDetails)

}

func PostHouseDetailsHandler(w http.ResponseWriter, r *http.Request) {
	var HouseDetails models.HouseDetails
	json.NewDecoder(r.Body).Decode(&HouseDetails)

	createdHouseDetails := db.DB.Create(&HouseDetails)
	err := createdHouseDetails.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&HouseDetails)
}

func UpdateHouseDetailsHandler(w http.ResponseWriter, r *http.Request) {
	var HouseDetails models.HouseDetails
	params := mux.Vars(r)

	db.DB.First(&HouseDetails, params["id"])

	json.NewDecoder(r.Body).Decode(&HouseDetails)

	if HouseDetails.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("House not found"))
		return
	}

	db.DB.Save(&HouseDetails)
	w.WriteHeader(http.StatusNoContent)
}

func DeleteHouseDetailsHandler(w http.ResponseWriter, r *http.Request) {

	var HouseDetails models.HouseDetails
	params := mux.Vars(r)

	db.DB.First(&HouseDetails, params["id"])

	if HouseDetails.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("House Details Not Found"))
		return
	}

	db.DB.Delete(&HouseDetails)
	w.WriteHeader(http.StatusNoContent)

}
