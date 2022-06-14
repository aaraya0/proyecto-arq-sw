package controllers

import (
	"net/http"
	"strconv"

	"github.com/aaraya0/arq-software/Integrador1/dto"
	service "github.com/aaraya0/arq-software/Integrador1/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetOrderById(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	log.Debug("Order id: " + c.Param("id"))

	var orderDto dto.OrderDto

	c.JSON(http.StatusOK, orderDto)
}

func GetOrdersByUId(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	log.Debug("Id: " + c.Param("id"))

	var ordersDto dto.OrdersDto
	id, _ := strconv.Atoi(c.Param("id"))
	ordersDto, err := service.OrderService.GetOrdersByUId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, ordersDto)
}

func OrderInsert(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var orderDtoInsert dto.OrderDtoInsert
	var orderDtoResp dto.OrderDtoResp
	err := c.BindJSON(&orderDtoInsert)

	log.Debug(orderDtoInsert)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	orderDtoResp, er := service.OrderService.AddOrder(orderDtoInsert)
	if er != nil {
		log.Error(er.Error())
		c.JSON(er.Status(), er.Error())
		return
	}
	for i := 0; i < len(orderDtoInsert.OrderDetails); i++ {
		_, e := service.DetailService.AddOrderDetail(orderDtoInsert.OrderDetails[i], orderDtoResp.Id)
		if e != nil {
			log.Error(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	c.JSON(http.StatusCreated, orderDtoInsert)
}
