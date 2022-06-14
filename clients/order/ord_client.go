package order

import (
	model "github.com/aaraya0/arq-software/Integrador1/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetOrderById(id int) model.Order {
	var order model.Order
	Db.Where("id = ?", id).First(&order)
	log.Debug("Order:", order)
	return order
}

func GetOrdersByUId(id int) model.Orders {
	var orders model.Orders
	Db.Where("user_id = ?", id).Find(&orders)
	log.Debug("Orders:", orders)
	return orders
}

func AddOrder(order model.Order) model.Order {

	result := Db.Create(&order)

	if result.Error != nil {

		order.Id = 0
		return order
	}
	log.Debug("Order Created: ", order.Id)
	return order
}
