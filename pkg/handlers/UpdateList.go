package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"todocrud/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) UpdateList(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedList models.List
	json.Unmarshal(body, &updatedList)

	var list models.List

	if result := h.DB.First(&list, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	list.Title = updatedList.Title
	list.Auther = updatedList.Auther
	list.Desc = updatedList.Desc

	h.DB.Save(&list)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Updated")

}
