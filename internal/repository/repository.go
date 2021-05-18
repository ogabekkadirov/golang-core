package repository

import (
	"github.com/jinzhu/gorm"
)


type Repositories struct {
	BaseRepo BaseRepo
	Users Users

}

func NewRepositories(db *gorm.DB) *Repositories{
	return &Repositories{
		BaseRepo: BaseRepo{},
		Users: NewUsersRepo(db),
	}
}