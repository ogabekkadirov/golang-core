package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)


func Check() gin.HandlerFunc{
	return func(c *gin.Context) {
		authorization := string(c.Request.Header.Get("Authorization"))

		if authorization != ""{
			auth := strings.SplitN(authorization, " ", 2)
			if len(auth) == 2 && auth[0] == "Basic" {
				// err := Auth
			}
		}
	}
}	