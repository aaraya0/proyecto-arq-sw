package model

type Product struct {
	Id            int     `gorm:"primaryKey"`
	Title         string  `gorm:"type:varchar(350);not null"`
	Author        string  `gorm:"type:varchar(350);not null"`
	UniversalCode string  `gorm:"type:varchar(350);not null"`
	Price         float32 `gorm:"type:varchar(350);not null"`
	Stock         int     `gorm:"type:varchar(350);not null"`

	//Picture       string  //`gorm:"picture_url"`
}

type Products []Product
