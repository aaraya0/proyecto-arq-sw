package controllers

import (
	"net/http"

	"github.com/aaraya0/arq-software/Integrador1/dto"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetDetailById(c *gin.Context) {
	log.Debug("Detail id: " + c.Param("id"))

	var orderDetailDto dto.DetailDto
	c.JSON(http.StatusOK, orderDetailDto)
}

func DetailInsert(c *gin.Context) {
	var detailDto dto.DetailDto
	err := c.BindJSON(&detailDto)

	log.Debug(detailDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, detailDto)
}
