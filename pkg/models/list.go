package models

type List struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Auther string `json:"auther"`
	Desc   string `json:"desc"`
}
