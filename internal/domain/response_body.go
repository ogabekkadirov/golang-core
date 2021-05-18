package domain

type ResponseBody struct {
	Pagination Pagination  `json:"pagination"`
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
}

func NewResponseBody(pagination Pagination, data interface{}, total int64) ResponseBody {

	return ResponseBody{
		Pagination: pagination,
		Data:       data,
		Total:      total,
	}
}