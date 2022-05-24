package model

type Order struct {
	Id int `gorm:"type:int;not null"`
	Date  int `gorm:"type:int;not null"`
	Detail OrderDets
	Total int // como se hace la suma?
}

type Orders []Order
