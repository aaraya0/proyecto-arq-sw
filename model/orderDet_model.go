package model

type OrderDet struct {
	Amount  float32 `gorm:"type:float;not null"`
	Product []Product
	Price   float32 `gorm:"type:float;not null"`
	Id      int     `gorm:"type:int(11);"`
}

type OrderDets []OrderDet
