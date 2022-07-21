package detail

import (
	model "mvc/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InsertOrderDetail(orderDetail model.Detail) model.Detail {
	result := Db.Create(&orderDetail)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Detail Created: ", orderDetail.Id)
	result1 := Db.Model(&model.Product{}).Where("id = ?", orderDetail.Product_Id).Update("stock", gorm.Expr("stock - ? ", orderDetail.Quantity))

	if result1.Error != nil {
		log.Error("Producto no encontrado")
	}

	return orderDetail
}

func GetOrderDetailByOrderId(orderId int) model.Details {
	var orderDetails model.Details

	Db.Where("order_id = ?", orderId).Find(&orderDetails)
	log.Debug("Detail: ", orderDetails)

	return orderDetails
}

func InsertOrderDetails(orderDetails model.Details) model.Details {

	for _, orderDetail := range orderDetails {

		result := Db.Create(&orderDetail)

		if result.Error != nil {
			log.Error("Error al crear")
		}
		log.Debug("Detail Created: ", orderDetail.Id)
		result1 := Db.Model(&model.Product{}).Where("id = ?", orderDetail.Product_Id).Update("stock", gorm.Expr("stock - ? ", orderDetail.Quantity))

		if result1.Error != nil {
			log.Error("Producto no encontrado")
		}
	}

	return orderDetails
}
