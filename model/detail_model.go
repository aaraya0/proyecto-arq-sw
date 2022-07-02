package model

type Detail struct {
	Id         int     `gorm:"primaryKey"`
	UnitPrice  float32 `gorm:"type:float;not null"`
	Quantity   float32 `gorm:"type:int;not null"`
	Total      float32 `gorm:"type:float;not null"`
	Name       string  `gorm:"type:varchar(550);not null"`
	Product_Id int     `gorm:"type:int;not null"`
	Order_Id   int     `gorm:"type:int;not null"`
}

type Details []Detail
