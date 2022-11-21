package repository

import (
	"golang-core/internal/domain"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)




type Users interface{
	GetList(domain.Pagination)([]domain.User, error)
	GetById(int64)( domain.User, error)
	Create(*domain.User) error
	Update(domain.User,  *domain.User) error
	Delete(domain.User) error
}

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) *UsersRepo{
	return &UsersRepo{db:db}
}

func (repo *UsersRepo) GetList(pagination domain.Pagination)(result []domain.User,  err error){
	

	err = repo.db.Select(&result, 
		`SELECT id,
						username,
						fullname,
						status_id,
						password 
					FROM users 
					ORDER BY id ASC 
					OFFSET $1 LIMIT $2`,pagination.Limit,pagination.Page)

	return 
}

func (repo *UsersRepo) GetById(id int64)(result domain.User, err error){
	
	err = repo.db.Get(&result,"SELECT * FROM users WHERE id=$1",id)
	
	if err  != nil{
		return
	}

	return 
}
func (repo *UsersRepo) Create(user *domain.User)(err error){

	_,err = repo.db.NamedExec(`INSERT INTO person 
																(username, 
																	fullname, 
																	status_id, 
																	password) 
														 VALUES 
																(:username,
																	:fullname, 
																	:status_id, 
																	:password)`,user)

	return
}
func (repo *UsersRepo) Update(input domain.User, model *domain.User)(err error){
	

	_,err = repo.db.NamedExec(`UPDATE users 
															SET	
																(username, 
																	fullname, 
																	status_id, 
																	password) 
														  VALUES 
																(:username,
																	:fullname, 
																	:status_id, 
																	:password)`,input)
	return
}

func (repo *UsersRepo) Delete(model domain.User)(err error){

	_,err = repo.db.Exec("DELETE FROM Users WHERE user_id=$1", model.ID)

	return
}