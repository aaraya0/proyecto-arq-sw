package services

import (
	userCliente "github.com/aaraya0/arq-software/Integrador1/mvc/clients/user"
	"github.com/aaraya0/arq-software/Integrador1/mvc/dto"
	"github.com/aaraya0/arq-software/Integrador1/mvc/model"
	e "github.com/aaraya0/arq-software/Integrador1/mvc/utils/errors"
)

type userService struct{}

type userServiceInterface interface {
	GetUserByUname(uname string) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserByUname(uname string) (dto.UserDto, e.ApiError) {

	var user model.User = userCliente.GetUserByUname(uname)
	var userDto dto.UserDto

	if user.UserName == "" {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Id = user.Id
	userDto.Password = user.Password
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.UserName = user.UserName
		userDto.Id = user.Id
		userDto.Password = user.Password
		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}
