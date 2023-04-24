package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixVNolasco/real-state-golang/db"
	"github.com/FelixVNolasco/real-state-golang/models"
	"github.com/gorilla/mux"
)

func GetHouseGalleryHandler(w http.ResponseWriter, r *http.Request) {

	var HouseGallery models.HouseGallery
	params := mux.Vars(r)

	db.DB.First(&HouseGallery, params["id"])

	if HouseGallery.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("House Gallery Not Found"))
		return
	}

	db.DB.Model(&HouseGallery).Association("Photos").Find(&HouseGallery.Photos)

	json.NewEncoder(w).Encode(&HouseGallery)

}

func PostHouseGalleryHandler(w http.ResponseWriter, r *http.Request) {
	var HouseGallery models.HouseGallery
	json.NewDecoder(r.Body).Decode(&HouseGallery)

	createdHouseGallery := db.DB.Create(&HouseGallery)
	err := createdHouseGallery.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&createdHouseGallery)
}
