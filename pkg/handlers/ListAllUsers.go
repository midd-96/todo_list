package handlers

import (
	"encoding/json"
	"net/http"
	"todocrud/pkg/models"
)

func (h handler) ListAllUsers(w http.ResponseWriter, r *http.Request) {

	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		json.NewEncoder(w).Encode(result.Error)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
