package app

import (
	categoryController "github.com/aaraya0/arq-software/Integrador1/controllers/category"
	detailController "github.com/aaraya0/arq-software/Integrador1/controllers/detail"
	orderController "github.com/aaraya0/arq-software/Integrador1/controllers/order"
	productController "github.com/aaraya0/arq-software/Integrador1/controllers/product"
	userController "github.com/aaraya0/arq-software/Integrador1/controllers/user"
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

	log.Info("Finishing mappings configurations")

}
