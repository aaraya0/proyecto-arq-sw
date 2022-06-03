package model

type Address struct {
	Id      int    `gorm:"primaryKey"`
	Street  string `gorm:"type:varchar(350);not null"`
	Number  int    `gorm:"type:int;not null"`
	Country string `gorm:"type:varchar(350);not null"`
	City    string `gorm:"type:varchar(350);not null"`
	CP      int    `gorm:"type:int;not null"`
}

type Addresses []Address
