package utils

import (
	"errors"
	"golang-core/internal/domain"
	"golang-core/utils/cerror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdParamFromRequest(ctx *gin.Context)(id int64, cError domain.Error){
	
	idParam := ctx.Param("id")

	if idParam == "" {
		err := errors.New("id not found")
		cError = cerror.NewError(http.StatusBadRequest, err)
		return
	}

	id, err := strconv.ParseInt(idParam, 10,64)

	if err != nil {
		err := errors.New("id is of invalid type")
		cError = cerror.NewError(http.StatusBadRequest, err)
		return
	}
	
	return
}