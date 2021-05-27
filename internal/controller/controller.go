package controller

import "golang-core/internal/service"

type Controllers struct {
	Users Users
}

type Deps struct {
	Services *service.Sevices
}

func NewControllers(deps Deps) *Controllers{
	return &Controllers{
		Users: NewUsersController(deps.Services.Users),
	}
}