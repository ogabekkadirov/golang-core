package repository

import (
	"golang-core/internal/domain"

	"github.com/jinzhu/gorm"
)


type BaseRepo struct {
}

func (repo BaseRepo) Paginate(dbQuery *gorm.DB, pagination domain.Pagination)(result *gorm.DB){

	offset := pagination.Limit*(pagination.Page-1)

	result = dbQuery.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	return
}

func (repo BaseRepo) With(dbQuery *gorm.DB, loads []string)(result *gorm.DB){

	result = dbQuery
	
	for _,s := range loads{
		result = result.Preload(s)
	}

	return
}

// func (repo BaseRepo) WithCondition(dbQuery *gorm.DB, loads map[string]interface{})(result *gorm.DB){

// 	result = dbQuery
	

// 	return
// }