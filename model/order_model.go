package model

type Order struct {
	Id      int     `gorm:"type:int;not null"`
	User_id int     `gorm:"type:int;not null"`
	Total   float32 `gorm:"type:float;not null"`
}

type Orders []Order
