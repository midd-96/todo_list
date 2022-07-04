package handlers

import (
	// "crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"todocrud/pkg/models"
)

func (h handler) AddList(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var list models.List
	json.Unmarshal(body, &list)

	// list.Id = rand.Intn(100)
	// mocks.Lists = append(mocks.Lists, list)

	if result := h.DB.Create(&list); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
