package http

import (
	"golang-core/internal/controller"
	"golang-core/pkg/auth"
)


type Handler struct {
	tokenManager auth.TokenManager
	userController controller.UserController
	controller *controller.Controllers
}

func NewHandler(cont *controller.Controllers) *Handler {
	return &Handler{
		controller: cont,
	}
}