package models

import "gorm.io/gorm"

type Departments struct {
	gorm.Model

	Dep_Id string `json:"dep_id" gorm:"primary_key"`
	Name   string `json:"dep_name" gorm:"not null;unique"`
}
