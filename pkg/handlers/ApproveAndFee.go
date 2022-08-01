package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"todocrud/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) AppoveAndFee(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	email, _ := vars["email"]

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedDoctor models.Doctor
	json.Unmarshal(body, &updatedDoctor)

	var doctor models.Doctor

	result := h.DB.First(&doctor, email)
	if result == nil {
		json.NewEncoder(w).Encode("Invalid Email id")
		json.NewEncoder(w).Encode(result.Error)
	} else {
		// doctor.Approvel = updatedDoctor.Approvel
		// doctor.Fee = updatedDoctor.Fee
		// h.DB.Save(&doctor)
		//h.DB.Model(&doctor).Select("approvel", "fee").Updates(models.Doctor{Approvel: updatedDoctor.Approvel, Fee: updatedDoctor.Fee})
		// h.DB.Model(&models.Doctor{Email: updatedDoctor.Email}).Updates(models.Doctor{
		// 	Approvel: updatedDoctor.Approvel,
		// 	Fee:      updatedDoctor.Fee,
		// })

		if err := h.DB.Where(models.Doctor{Email: updatedDoctor.Email}).
			Assign(models.Doctor{Email: updatedDoctor.Email, Approvel: updatedDoctor.Approvel, Fee: updatedDoctor.Fee}).
			FirstOrCreate(&models.Doctor{}).Error; err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Updated")
	}

	// list.Title = updatedList.Title
	// list.Auther = updatedList.Auther
	// list.Desc = updatedList.Desc

	//h.DB.Save(&doctor)

}
