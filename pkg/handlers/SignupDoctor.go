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

func (h handler) SignupDoctor(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var doctor models.Doctor

	json.Unmarshal(body, &doctor)

	hashedPassword, err := util.HashPassword(doctor.Password)
	if err != nil {
		json.NewEncoder(w).Encode("Internal server error can't convert to hash password")
		return
	}
	doctor.Password = hashedPassword

	valid_email := util.ValidateEmail(doctor.Email)
	if valid_email == false {
		json.NewEncoder(w).Encode("Email not valid")
		return
	}

	valid_phone := util.ValidatePhone(doctor.Phone)
	if valid_phone == false {
		json.NewEncoder(w).Encode("Phone number is not valid")
		return
	}

	result := h.DB.Create(&doctor)

	if result.Error != nil {
		json.NewEncoder(w).Encode("Can't signup")
		json.NewEncoder(w).Encode(result.Error)

	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Created")
		json.NewEncoder(w).Encode(doctor.Last_name)

	}

}
