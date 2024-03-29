package v1

import (
	"golang-core/internal/controller"
	"golang-core/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Handler struct {
	Controllers *controller.Controllers
}

func NewHandler (conts *controller.Controllers) *Handler{
	return &Handler{
		Controllers: conts,
	}
}

func (h *Handler) Init(api *gin.RouterGroup){
	v1 := api.Group("/v1")
	{
		h.InitUsersRoutes(v1)
		v1.GET("/test", func(ctx *gin.Context) {
			response.SuccessResult(ctx, http.StatusOK, "v1 routes are working :).")
		})
	}
}