package domain

import "golang.org/x/crypto/bcrypt"

type CUUser struct {
	Username string `json:"username" form:"username" binding:"min=3,max=30,required"`
	Fullname string `json:"fullname" form:"fullname" binding:"gte=3,lte=150,required"`
	Password []byte `json:"-" form:"-" binding:"min=6,required"`
}

type User struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Fullname string `json:"fullname" form:"fullname"`
	Password []byte `json:"-" form:"-"`
}

func (u *User) SetPassword(password string) {
	u.Password,_= bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}