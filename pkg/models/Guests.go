package models

type Guests struct {
	Id         int         `json:"id" gorm:"primary_key"`
	Name       string      `json:"dep_name" gorm:"not null"`
	First_name string      `json:"first_name" gorm:"not null"`
	Last_name  string      `json:"last_name" gorm:"not null"`
	Email      string      `json:"email" gorm:"unique;not null" valid:"email"`
	Phone      string      `json:"phone" gorm:"not null"`
	Department Departments `gorm:"foreignKey:Name" `
	Doc_Id     Doctor      `gorm:"foreignKey:Id"`
}
