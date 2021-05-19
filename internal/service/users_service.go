package service

import (
	"errors"
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
	// Create(user domain.CUUser)(result domain.User, cError domain.Error)
}

type UsersService struct {
	repo repository.Users
	baseRepo repository.BaseRepo
}

func NewUsersService(repo repository.Users) *UsersService {

	return &UsersService{
		repo: repo,
	}

}

func (s *UsersService) GetAll(pagination domain.Pagination)(result domain.ResponseBody, cError domain.Error){

	data, err := s.repo.GetTable()

	if err != nil {
		cError = cerror.NewError(http.StatusInternalServerError,err)
		return 
	}

	data.Count(&total)

	data = s.baseRepo.Paginate(data, pagination)

	result = domain.NewResponseBody(pagination, data, total)

	return
}

func (s *UsersService) GetById(id int64)(result domain.User, cError domain.Error){

	result, err = s.repo.GetById(id)

	if err != nil {
		cError = cerror.NewError(http.StatusInternalServerError, err)
		return
	}

	if result.ID == 0 {
		err = errors.New("Data Not Found")
		cError = cerror.NewError(http.StatusNotFound, err)
	}
	return
}
// func (s *UsersService) Create(input domain.CUUser)(result domain.User, cError domain.Error){
// 	err = s.repo.Create(&input)
// }