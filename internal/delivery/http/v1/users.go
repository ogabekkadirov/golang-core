package v1

import "github.com/gin-gonic/gin"

func (h *Handler) InitUsersRoutes(api *gin.RouterGroup){

	UsersController := h.Controllers.Users

	users := api.Group("/users")
	{
		users.GET("/",UsersController.Index)
		users.POST("/",UsersController.Store)
		users.GET("/:id",UsersController.Get)
	}
}