package service

import (
	"golang-core/internal/repository"
)


type Sevices struct {
	Users Users
}

type Deps struct {
	Repos *repository.Repositories
}
func NewServices(deps Deps) *Sevices {
	return &Sevices{
		Users: NewUsersService(deps.Repos.Users),
	}
}