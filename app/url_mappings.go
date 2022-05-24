package app

import (
	productController "Integrador1/controllers/product"
	userController "Integrador1/controllers/user"
	_ cartController "Integrador1/controllers/cart"
	_ categoryController "Integrador1/controllers/category"
    _ orderController "Integrador1/controllers/order"
    _ orderDetController "Integrador1/controllers/order_detail"
    _ addressController "Integrador1/controllers/address"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/product/:product_id", productController.GetProductByEan)
	router.GET("/product", productController.GetProducts)

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert)


	log.Info("Finishing mappings configurations")

}
