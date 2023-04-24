package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixVNolasco/real-state-golang/db"
	"github.com/FelixVNolasco/real-state-golang/models"
	"github.com/gorilla/mux"
)

func GetPhotoHandler(w http.ResponseWriter, r *http.Request) {

	var photo models.Photo
	params := mux.Vars(r)

	db.DB.First(&photo, params["id"])

	if photo.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Photo Not Found"))
		return
	}

	json.NewEncoder(w).Encode(&photo)

}

func PostPhotoHandler(w http.ResponseWriter, r *http.Request) {
	var photo models.Photo
	json.NewDecoder(r.Body).Decode(&photo)

	createdPhoto := db.DB.Create(&photo)
	err := createdPhoto.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&createdPhoto)
}
