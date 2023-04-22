package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixVNolasco/real-state-golang/db"
	"github.com/FelixVNolasco/real-state-golang/models"
)

func GetAgentsHandler(w http.ResponseWriter, r *http.Request) {
	var agents []models.Agent

	db.DB.Find(&agents)
	json.NewEncoder(w).Encode(&agents)
}

func GetAgentHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get agent"))
}

func PostAgentsHandler(w http.ResponseWriter, r *http.Request) {

	var agent models.Agent

	json.NewDecoder(r.Body).Decode(&agent)

	createdAgent := db.DB.Create(&agent)

	err := createdAgent.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&agent)
}

func DeleteAgentHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
