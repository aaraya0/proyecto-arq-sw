package product

import (
	model "github.com/aaraya0/arq-software/Integrador1/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetProductById(id int) model.Product {
	var product model.Product
	Db.Where("id=?", id).First(&product)

	log.Debug("Product:", product)
	return product
}

func GetProducts() model.Products {
	var products model.Products
	Db.Find(&products)
	log.Debug("Products: ", products)
	return products

}

//InsertProduct
