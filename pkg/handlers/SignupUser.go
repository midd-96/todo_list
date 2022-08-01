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

func (h handler) SignupUser(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		json.NewEncoder(w).Encode("Internal server error can't convert to hash password")
		return
	}
	user.Password = hashedPassword

	valid_email := util.ValidateEmail(user.Email)
	if valid_email == false {
		json.NewEncoder(w).Encode("Email not valid")
		return
	}
	valid_phone := util.ValidatePhone(user.Phone)
	if valid_phone == false {
		json.NewEncoder(w).Encode("Phone number is not valid")
		return
	}

	result := h.DB.Create(&user)

	if result.Error != nil {
		json.NewEncoder(w).Encode("Can't signup")
		json.NewEncoder(w).Encode(result.Error)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Created")
		json.NewEncoder(w).Encode(user.Last_name)
	}

}
