package model

type Product struct {
	Id           int    `gorm:"primaryKey"`
	Product_Name string `gorm:"type:varchar(350);not null"`
	Cost         int    `gorm:"type:int;not null"`
}

type Products []Product
