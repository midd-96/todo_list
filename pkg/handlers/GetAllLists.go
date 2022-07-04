package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todocrud/pkg/models"
)

func (h handler) GetAllLists(w http.ResponseWriter, r *http.Request) {

	var lists []models.List

	if result := h.DB.Find(&lists); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lists)
}
