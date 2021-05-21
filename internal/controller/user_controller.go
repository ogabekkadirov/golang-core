package controller

import (
	"golang-core/internal/domain"
	"golang-core/internal/service"
	"golang-core/utils"
	"golang-core/utils/cerror"
	"golang-core/utils/pagination"
	"golang-core/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Users interface {
	Index(ctx *gin.Context)
	Get(ctx *gin.Context)
	Store(ctx *gin.Context)
}

type UserController struct {
	service service.Users
}

func NewUserController(service service.Users) *UserController {
	return &UserController{
		service: service,
	}
}

func (c *UserController) Index(ctx *gin.Context){

	pagination := pagination.GeneratePaginationFromRequest(ctx)

	result, err := c.service.GetAll(pagination)

	if err.Err != nil {
		response.ErrorResult(ctx, err)
		return 
	}

	response.SuccessResult(ctx, http.StatusOK, result)
	return
}

func (c *UserController) Get(ctx *gin.Context){

	id, err := utils.GetIdParamFromRequest(ctx)

	if err.Err != nil{
		response.ErrorResult(ctx, err)
		return
	}

	result, err := c.service.GetById(id)

	if err.Err != nil{
		response.ErrorResult(ctx, err)
		return
	}

	response.SuccessResult(ctx, http.StatusOK, result)
	
	return
}

func (c *UserController) Store(ctx *gin.Context){
	
	input := domain.CUUser{}

	if err := ctx.ShouldBindJSON(&input); err != nil{
		cError := cerror.NewError(http.StatusBadRequest, err)
		response.ErrorResult(ctx, cError)
		return
	}

	

}
