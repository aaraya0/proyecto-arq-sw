package controllers

import (
	"net/http"
	"strconv"

	"mvc/dto"
	service "mvc/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func OrderInsert(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "POST")
	var orderDto dto.OrderDto
	err := c.BindJSON(&orderDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	orderDto, er := service.OrderService.InsertOrder(orderDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, orderDto)
}

func GetOrdersByUserId(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var ordersDto dto.OrdersDto
	token, _ := strconv.Atoi(c.Param("userId"))
	ordersDto, err := service.OrderService.GetOrdersByUserId(token)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, ordersDto)
}
