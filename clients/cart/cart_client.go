package cart

import (
	model "github.com/aaraya0/arq-software/Integrador1/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func AddProduct(cart model.Cart) model.Cart {
	result := Db.Create(&cart)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Product Added: ", cart.Title)
	return cart
}

func GetCart() model.Carts {
	var carts model.Carts
	Db.Find(&carts)
	log.Debug("Cart: ", carts)
	return carts
}

//InsertProduct
