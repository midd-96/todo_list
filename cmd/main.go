package main

import (
	"log"
	"net/http"
	"todocrud/pkg/db"
	"todocrud/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {

	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/lists", h.GetAllLists).Methods(http.MethodGet)
	router.HandleFunc("/lists/{id}", h.GetList).Methods(http.MethodGet)
	router.HandleFunc("/lists", h.AddList).Methods(http.MethodPost)
	router.HandleFunc("/lists/{id}", h.UpdateList).Methods(http.MethodPut)
	router.HandleFunc("/lists/{id}", h.DeleteList).Methods(http.MethodDelete)

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}
