package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"todocrud/pkg/models"
)

func (h handler) AddDept(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var dept models.Departments
	json.Unmarshal(body, &dept)

	if result := h.DB.Create(&dept); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("DeptCreated")
	json.NewEncoder(w).Encode(dept)
}
