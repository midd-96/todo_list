package handlers

import (
	// "crypto/rand"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"todocrud/pkg/models"
	"todocrud/util"
)

func (h handler) GuestBooking(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var guest models.Guests
	json.Unmarshal(body, &guest)

	valid_email := util.ValidateEmail(guest.Email)
	if valid_email == false {
		json.NewEncoder(w).Encode("Email not valid")
		return
	}
	valid_phone := util.ValidatePhone(guest.Phone)
	if valid_phone == false {
		json.NewEncoder(w).Encode("Phone number is not valid")
		return
	}

	result := h.DB.Create(&guest)

	if result.Error != nil {
		json.NewEncoder(w).Encode("Can't complete")
		json.NewEncoder(w).Encode(result.Error)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Guest Details saved")
		json.NewEncoder(w).Encode(guest.Last_name)
	}

}
