package handlers

import (
	//"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todocrud/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) GetList(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var list models.User

	if result := h.DB.First(&list, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(list)

}
