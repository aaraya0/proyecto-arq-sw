package model

type Order struct {
	Id int `gorm:"type:int;not null"`

	OrderDetail []Detail
	User        []User
	Total       float32 `gorm:"type:float;not null"`
}

type Orders []Order
