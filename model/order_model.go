package model

import (
	"time"
)

type Order struct {
	Id          int       `gorm:"primaryKey"`
	TotalAmount float32   `gorm:"type:float;not null"`
	Date        time.Time `gorm:"not null"`
	User_Id     int       `gorm:"type:int;not null"`
}

type Orders []Order
