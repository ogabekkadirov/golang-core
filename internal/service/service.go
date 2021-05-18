package service

import "golang-core/internal/repository"

type Sevices struct {
	Users Users
}

type Deps struct {
	Repos *repository.Repositories
}

func NewService(deps Deps) *Sevices {
	return &Sevices{
		Users: NewUsersService(deps.Repos.Users),
	}
}