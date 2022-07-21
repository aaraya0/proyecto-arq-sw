package order

import (
	model "mvc/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InsertOrder(order model.Order) model.Order {
	result := Db.Create(&order)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Order Created: ", order.Id)
	return order
}

func GetOrdersByUserId(userId int) model.Orders {
	var orders model.Orders

	log.Debug("userId: ", userId)
	Db.Where("user_id= ?", userId).Find(&orders)
	log.Debug("Order: ", orders)

	return orders
}

func UpdateMontoFinal(monto float32, orderId int) float32 {
	result := Db.Model(&model.Order{}).Where("id = ?", orderId).Update("total_amount", monto)

	if result.Error != nil {
		log.Error("Order no encontrada")
	}
	return monto
}
