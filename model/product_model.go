package model

type Product struct {
	Id          int     `gorm:"type:int(11);"`
	Title       string  `gorm:"type:varchar(2500);"`
	Author      string  `gorm:"type:varchar(250);"`
	Image       string  `gorm:"type:varchar(250);"`
	Price       float32 `gorm:"type:float;"`
	Stock       int     `gorm:"type:int(11);"`
	Category_id int     `gorm:"foreignKey:category_id"`
}

type Products []Product
