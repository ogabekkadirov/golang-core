package repository

import "github.com/jmoiron/sqlx"


type Repositories struct {
	BaseRepo BaseRepo
	Users Users
}

func NewRepositories(db *sqlx.DB) *Repositories{
	return &Repositories{
		BaseRepo: BaseRepo{},
		Users: NewUsersRepo(db),
	}
}