package repository

import (
	"golang-core/internal/domain"

	"github.com/jinzhu/gorm"
)




type Users interface{
	All()(result []domain.User, err error)
}

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo{
	return &UsersRepo{db:db}
}

func (repo *UsersRepo) All()(result []domain.User, err error){
	
	repo.db.Table("users").Find(&result)
	
	return 
}