package model

type Detail struct {
	Quantity  int `gorm:"type:float;not null"`
	Product []Product `gorm:"type:int(11)"`
	Price   float32 `gorm:"type:float;not null"`
	Id      int     `gorm:"type:int(11);"`
}

type Details []Detail
