package pagination

import (
	"golang-core/config"
	"golang-core/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)



func GeneratePaginationFromRequest(ctx *gin.Context) domain.Pagination{
	
	limit := config.PAGINATION_LIMIT
	sort  := config.PAGINATION_SORT 
	page  := 1

	query := ctx.Request.URL.Query()

	for key, value := range query{
		queryValue := value[len(value)-1]

		switch key{

			case "limit":
				limit,_ = strconv.Atoi(queryValue)
				break

			case "page":
				page,_ = strconv.Atoi(queryValue)
				break

			case "sort":
				sort = queryValue
				break
			}
	}
	return domain.Pagination{
		Limit: limit,
		Page: page,
		Sort: sort,
	}
}