package v1

import "github.com/gin-gonic/gin"

func (h *Handler) InitUsersRoutes(api *gin.RouterGroup){
	users := api.Group("/users")
	{
		users.GET("",)
	}
}