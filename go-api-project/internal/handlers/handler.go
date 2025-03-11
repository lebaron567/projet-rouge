package handlers

import (
	"encoding/json"
	"go-api-project/internal/db"
	"go-api-project/internal/models"
	"net/http"
)

// HandleGet responds to GET requests with a simple message.
func HandleGet(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Bienvenue sur l'API Go !"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// HandlePost responds to POST requests with the received data.
func HandlePost(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{"message": "POST request successful", "data": data}
	json.NewEncoder(w).Encode(response)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
