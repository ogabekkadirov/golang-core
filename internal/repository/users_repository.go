package repository

import (
	"golang-core/internal/domain"

	"github.com/jinzhu/gorm"
)




type Users interface{
	GetTable()(result *gorm.DB, err error)
	GetById(id int64)(result domain.User, err error)
	Create(user *domain.User) error
	Update(input domain.User, model *domain.User) error
	Delete(model domain.User) error
}

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo{
	return &UsersRepo{db:db}
}

func (repo *UsersRepo) GetTable()(result *gorm.DB,  err error){
	
	result = repo.db.Table("users")
	
	if result.Error != nil{
		err = result.Error
	}

	return 
}

func (repo *UsersRepo) GetById(id int64)(result domain.User, err error){
	
	data := repo.db.Table("users").First(&result, id)
	
	if data.Error != nil{
		err = data.Error
		return
	}

	return 
}
func (repo *UsersRepo) Create(user *domain.User)(err error){

	result := repo.db.Create(&user)
	if result.Error != nil {
		err = result.Error
	}

	return
}
func (repo *UsersRepo) Update(input domain.User, model *domain.User)(err error){
	
	result := repo.db.Model(&model).Omit("CreatedAt", "Password").Update(input)

	if result.Error != nil {
		err = result.Error
	}

	return
}

func (repo *UsersRepo) Delete(model domain.User)(err error){

	result := repo.db.Delete(model)

	if result.Error != nil{
		err = result.Error
	}
	return
}