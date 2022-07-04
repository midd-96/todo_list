package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todocrud/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) DeleteList(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var list models.List

	if result := h.DB.First(&list, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&list)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
