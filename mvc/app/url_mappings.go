package app

import (
	addrController "github.com/aaraya0/arq-software/Integrador1/mvc/controllers/address"
	categoryController "github.com/aaraya0/arq-software/Integrador1/mvc/controllers/category"
	detailController "github.com/aaraya0/arq-software/Integrador1/mvc/controllers/detail"
	orderController "github.com/aaraya0/arq-software/Integrador1/mvc/controllers/order"
	productController "github.com/aaraya0/arq-software/Integrador1/mvc/controllers/product"
	userController "github.com/aaraya0/arq-software/Integrador1/mvc/controllers/user"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/product", productController.GetProducts)
	router.GET("/products/:category_id", productController.GetProductsByCategoryId)

	// Users Mapping
	router.GET("/user/:uname", userController.GetUserByUname)
	router.GET("/user", userController.GetUsers)

	// Category mapping
	router.GET("/category/:id", categoryController.GetCategoryById)
	router.GET("/category", categoryController.GetCategories)

	// OrderDetail mapping
	router.GET("/orderDetailOrder/:orderId", detailController.GetOrderDetailByOrderId)

	//Order mapping
	router.POST("/order", orderController.OrderInsert)
	router.GET("/order/:userId", orderController.GetOrdersByUserId)

	//Address mapping
	router.POST("/address", addrController.AddressInsert)
	router.GET("/address/:userId", addrController.GetAddrByUserId)

	log.Info("Finishing mappings configurations")

}
