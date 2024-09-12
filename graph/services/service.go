package services

import (
	"context"

	"github.com/Nishi05/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type RepositoryService interface {
	GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error)
}

type Services interface {
	UserService
	RepositoryService
}

type services struct {
	*userService
	*repositoryService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		&userService{
			exec: exec,
		},
		&repositoryService{
			exec: exec,
		},
	}
}
