package models

import (
	"time"

	_ "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

//separate each model into separate files..
type User struct {
	gorm.Model

	Id               int        `json:"id" gorm:"primary_key"`
	First_name       string     `json:"first_name" gorm:"not null"`
	Last_name        string     `json:"last_name" gorm:"not null"`
	Email            string     `json:"email" gorm:"unique;not null" validate:"email"`
	Phone            string     `json:"phone" gorm:"unique;not null"`
	Password         string     `json:"password" gorm:"not null" validate:"min=6)"`
	Last_appointment time.Month `json:"last_appointment"`
	Role             int        `json:"role"`
}
