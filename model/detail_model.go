package model

type Detail struct {
	Quantity     int     `gorm:"type:float;not null"`
	Product_id   int     `gorm:"type:int(11)"`
	Price        float32 `gorm:"type:float;not null"`
	Id           int     `gorm:"type:int(11);"`
	Order_id     int     `gorm:"type:int(11);"`
	Product_name string  `gorm:"type:varchar(250);"`
}

type Details []Detail
