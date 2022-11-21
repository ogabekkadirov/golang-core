package repository

import (
	"golang-core/internal/domain"

	"github.com/jmoiron/sqlx"
)


type BaseRepo struct {
}

func (repo BaseRepo) Paginate(dbQuery *sqlx.DB, pagination domain.Pagination)(result *sqlx.DB){

	// offset := pagination.Limit*(pagination.Page-1)

	// result = dbQuery.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	return
}

func (repo BaseRepo) With(dbQuery *sqlx.DB, loads []string)(result *sqlx.DB){

	// result = dbQuery
	
	// for _,s := range loads{
	// 	result = result.Preload(s)
	// }

	return
}

// func (repo BaseRepo) WithCondition(dbQuery *sqlx.DB, loads map[string]interface{})(result *sqlx.DB){

// 	result = dbQuery
	

// 	return
// }