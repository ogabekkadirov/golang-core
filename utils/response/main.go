package response

import (
	"golang-core/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)


func SuccessResult(ctx *gin.Context, httpStatus int, result interface{})(mapResult map[string]interface{}){

	mapResult = make(map[string]interface{})
	mapResult["success"] = true
	mapResult["message"] = "ok"
	mapResult["result"] = result
	mapResult["status"] = httpStatus
	
	ctx.JSON(httpStatus, mapResult)

	return
}

func ErrorResult(ctx *gin.Context, err domain.Error)(mapResult map[string]interface{}){

	mapResult = make(map[string]interface{})
	mapResult["success"] = false
	mapResult["message"] = http.StatusText(err.Code)
	mapResult["status"] = err.Code
	mapResult["error_message"] = err.Err.Error()

	ctx.AbortWithStatusJSON(err.Code, mapResult)

	return
}