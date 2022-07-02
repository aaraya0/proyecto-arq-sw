package controllers

import (
	"net/http"
	"strconv"

	"github.com/aaraya0/arq-software/Integrador1/dto"
	"github.com/aaraya0/arq-software/Integrador1/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetOrderDetailByOrderId(c *gin.Context) {
	log.Debug("Detail id to load: " + c.Param("orderId"))

	orderId, _ := strconv.Atoi(c.Param("orderId"))
	var orderDetailResDto dto.DetailsDto

	orderDetailResDto, err := services.OrderDetailService.GetOrderDetailByOrderId(orderId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orderDetailResDto)
}
