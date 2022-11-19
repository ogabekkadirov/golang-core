package http

import (
	"golang-core/internal/controller"
	"golang-core/internal/delivery/http/middlewares"
	v1 "golang-core/internal/delivery/http/v1"
	"golang-core/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Handler struct {
	controllers *controller.Controllers
}

func NewHandler(conts *controller.Controllers) *Handler {
	return &Handler{
		controllers: conts,
	}
}
func (h *Handler) Init(host, port string) *gin.Engine {

	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		middlewares.CorsMiddleware(),
	)

	router.GET("/ping", func(ctx *gin.Context) {
		response.SuccessResult(ctx, http.StatusOK, "pong")
	})

	h.InitApi(router)

	return router
}

func (h *Handler) InitApi(router *gin.Engine){

	handlerV1 := v1.NewHandler(h.controllers)

	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}

}