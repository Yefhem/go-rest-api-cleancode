package controller

import "github.com/Yefhem/rest-api-cleancode/service"

type UserController interface {
}

type userController struct {
	userService service.UserService
}
