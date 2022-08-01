package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model

	Id         int    `json:"id" gorm:"primary_key"`
	First_name string `json:"first_name" gorm:"not null"`
	Last_name  string `json:"last_name" gorm:"not null"`
	Email      string `json:"email" gorm:"unique;not null" valid:"email"`
	Phone      string `json:"phone" gorm:"unique;not null"`
	Password   string `json:"password" gorm:"not null" valid:"length(6|20)"`
	//Dep_code       string `json:"dep_code"`
	Department     string `json:"department" gorm:"not null"`
	Specialization string `json:"specialization"`
	Approvel       bool   `json:"approvel" gorm:"default:false"`
	Fee            int    `json:"fee"`
}
