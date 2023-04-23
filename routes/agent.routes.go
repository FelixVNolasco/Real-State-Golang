package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixVNolasco/real-state-golang/db"
	"github.com/FelixVNolasco/real-state-golang/models"
	"github.com/gorilla/mux"
)

func GetAgentsHandler(w http.ResponseWriter, r *http.Request) {
	var agents []models.Agent

	db.DB.Find(&agents)
	json.NewEncoder(w).Encode(&agents)
}

func GetAgentHandler(w http.ResponseWriter, r *http.Request) {

	var agent models.Agent
	params := mux.Vars(r)

	db.DB.First(&agent, params["id"])

	if agent.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&agent).Association("Houses").Find(&agent.Houses)

	json.NewEncoder(w).Encode(&agent)
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
	var agent models.Agent
	params := mux.Vars(r)

	db.DB.First(&agent, params["id"])

	if agent.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Delete(&agent)
	w.WriteHeader(http.StatusNoContent)
}
