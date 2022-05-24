package model

type Adress struct {
	Id       int    `gorm:"primaryKey"`
	StName   string `gorm:"type:varchar(350);not null"`
	StNumber int    `gorm:"type:int;not null"`
}

type Adresses []Adress
