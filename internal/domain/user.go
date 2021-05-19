package domain

type CUUser struct {
	Username string `json:"username";form:"username"`
	Fullname string `json:"fullname";form:"fullname"`
	Password []byte `json:"-";form:"-"`
}

type User struct {
	ID       uint   `json:"id";form:"id"`
	Username string `json:"username";form:"username"`
	Fullname string `json:"fullname";form:"fullname"`
	Password []byte `json:"-";form:"-"`
}