package service

import (
	"errors"
	"fmt"
	"golang-core/internal/domain"
	"golang-core/internal/repository"
	"golang-core/utils/cerror"
	"net/http"
)
var total int64
var err error

type Users interface {
	GetAll(pagination domain.Pagination)(result domain.ResponseBody, cError domain.Error)
	GetById(id int64)(result domain.User, cError domain.Error)
	Create(input domain.CUUser)(result domain.User, cError domain.Error)
}

type UsersService struct {
	Repo repository.Users
	BaseRepo repository.BaseRepo
}

func NewUsersService(repo repository.Users) *UsersService {

	return &UsersService{
		Repo: repo,
	}

}

func (s *UsersService) GetAll(pagination domain.Pagination)(result domain.ResponseBody, cError domain.Error){

	data, err := s.Repo.GetTable()

	if err != nil {
		cError = cerror.NewError(http.StatusInternalServerError,err)
		return 
	}

	data.Count(&total)

	list := []domain.User{}

	s.BaseRepo.Paginate(data, pagination).Scan(&list)

	result = domain.NewResponseBody(pagination, list, total)

	return
}

func (s *UsersService) GetById(id int64)(result domain.User, cError domain.Error){

	result, err = s.Repo.GetById(id)

	if err != nil {
		
		if err.Error() == "record not found" {
			err = errors.New("Data Not Found")
			cError = cerror.NewError(http.StatusNotFound, err)
			return
		}
		
		cError = cerror.NewError(http.StatusInternalServerError, err)
		return
	}
	return
}

func (s *UsersService) Create(input domain.CUUser)(result domain.User, cError domain.Error){

	result = domain.User{
		Username: input.Username,
		Fullname: input.Fullname,
		StatusId: 1,
	}

	result.SetPassword(string(input.Password))

	fmt.Println(result)
	
	err = s.Repo.Create(&result)

	if err != nil {
		cError = cerror.NewError(http.StatusInternalServerError, err)
	}

	return
}