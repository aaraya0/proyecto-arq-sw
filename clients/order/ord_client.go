package order

import (
	model "github.com/aaraya0/arq-software/Integrador1/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetOrderById(id int) model.Order {
	var order model.Order
	Db.Where("id=?", id).First(&order)

	log.Debug("Order:", order)
	return order
}

func GetOrders() model.Orders {
	var orders model.Orders
	Db.Find(&orders)
	log.Debug("Orders: ", orders)
	return orders

}