package domain

import (
	"golang.org/x/crypto/bcrypt"
)


type CUser struct {
	Username string `json:"username" form:"username" binding:"min=3,max=30,required"`
	Fullname string `json:"fullname" form:"fullname" binding:"min=3,max=150,required"`
	Password string `json:"password" form:"password" binding:"min=6,required"`
}

type User struct {
	ID       uint   `json:"id" form:"id" binding:"min=3,max=30,required"`
	Username string `json:"username" form:"username" binding:"min=3,max=150,required"`
	Fullname string `json:"fullname" form:"fullname"`
	StatusId int8   `json:"status_id" form:"status_id"`
	Password string `json:"-" form:"-"`
}

func (u *User) SetPassword(password string) {
	
	pswd,_:= bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.Password = string(pswd)
}