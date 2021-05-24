package v1

import (
	"golang-core/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Handler struct {
	controllers *controller.Controllers
}

func NewHandler (conts *controller.Controllers) *Handler{
	return &Handler{
		controllers: conts,
	}
}

func (h *Handler) Init(api *gin.RouterGroup){
	v1 := api.Group("/v1")
	{
		h.InitUsersRoutes(api)
		v1.GET("/testing", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "v1 routes are working :).")
		})
	}
}