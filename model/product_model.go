package model

type Product struct {
	Id     int    `gorm:"type:int(11);"`
	Title  string `gorm:"type:varchar(2500);"`
	Author string `gorm:"type:varchar(250);"`
	//UniversalCode string  `gorm:"type:varchar(100);"`
	Price float32 `gorm:"type:float;"`
	Stock int     `gorm:"type:int(11);"`

	//Picture       string  //`gorm:"picture_url"`
}

type Products []Product
