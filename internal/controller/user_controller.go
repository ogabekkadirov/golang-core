package controller

import (
	"golang-core/internal/domain"
	"golang-core/internal/service"
	"golang-core/utils/pagination"
	"golang-core/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Users interface {
	GetAll(pagination domain.Pagination)(result domain.ResponseBody, err error)
	GetById(id int64)(result domain.User, err error)
}

type UserController struct {
	service service.Users
}

func (c *UserController) Index(ctx *gin.Context){

	pagination := pagination.GeneratePaginationFromRequest(ctx)

	result, err := c.service.GetAll(pagination)

	if err.Err != nil {
		response.ErrorResult(ctx, http.StatusBadRequest,"")
		return 
	}

	response.SuccessResult(ctx, http.StatusOK, result)
	return
}

