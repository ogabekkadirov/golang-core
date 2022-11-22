package domain

type Pagination struct {
	Limit  int    `json:"limit";form:"limit"`
	Page   int    `json:"page";form:"page"`
	Sort   string `json:"sort";form:"sort"`
	Offset int    `json:"offset";form:"offset"`
}