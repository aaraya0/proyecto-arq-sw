package app

import (
	addrController "mvc/controllers/address"
	categoryController "mvc/controllers/category"
	detailController "mvc/controllers/detail"
	orderController "mvc/controllers/order"
	productController "mvc/controllers/product"
	userController "mvc/controllers/user"
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
