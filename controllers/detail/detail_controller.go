package controllers

import (
	"net/http"
	"strconv"

	"github.com/aaraya0/arq-software/Integrador1/dto"
	service "github.com/aaraya0/arq-software/Integrador1/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetDetailrById(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	log.Debug("Detail id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	var detailDto dto.DetailDto
	detailDto, err := service.DetailService.GetDetailById(id)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, detailDto)
}

func GetDetails(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var detailsDto dto.DetailsDto
	detailsDto, err := service.DetailService.GetDetails()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, detailsDto)

}