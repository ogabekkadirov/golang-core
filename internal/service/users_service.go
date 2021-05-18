package service

import (
	"golang-core/internal/domain"
	"golang-core/internal/repository"
)
var total int64

type Users interface {
	GetAll(pagination domain.Pagination)(result domain.ResponseBody, err error)
}

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {

	return &UsersService{
		repo: repo,
	}

}

func (service *UsersService) GetAll(pagination domain.Pagination)(result domain.ResponseBody, err error){

	data, err := service.repo.All()

	if err != nil {
		return 
	}

	result = domain.NewResponseBody(pagination, data, total)

	return
}