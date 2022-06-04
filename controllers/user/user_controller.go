package controllers

import (
	"net/http"

	"github.com/aaraya0/arq-software/Integrador1/dto"
	service "github.com/aaraya0/arq-software/Integrador1/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserByUname(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	log.Debug("Username to load: " + c.Param("uname"))
	var uname = c.Param("uname")
	var userDto dto.UserDto
	userDto, err := service.UserService.GetUserByUname(uname)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var usersDto dto.UsersDto
	usersDto, err := service.UserService.GetUsers()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, usersDto)

}
