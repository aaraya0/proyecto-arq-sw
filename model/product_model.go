package model

type Product struct {
	Id     int    `gorm:"type:int(11);"`
	Title  string `gorm:"type:varchar(2500);"`
	Author string `gorm:"type:varchar(250);"`

	Price      float32 `gorm:"type:float;"`
	Stock      int     `gorm:"type:int(11);"`
	CategoryId int     `gorm:"foreignKey:id_category"`
}

type Products []Product
