package services

import (
	"context"

	"github.com/Nishi05/graphql-sample/graph/db"
	"github.com/Nishi05/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repositoryService struct {
	exec boil.ContextExecutor
}

func convertRepository(repo *db.Repository) *model.Repository {
	return &model.Repository{
		ID:        repo.ID,
		Name:      repo.Name,
		Owner:     &model.User{ID: repo.Owner},
		CreatedAt: repo.CreatedAt,
	}
}

func (r *repositoryService) GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error) {
	repo, err := db.Repositories(
		qm.Select(
			db.RepositoryColumns.ID,
			db.RepositoryColumns.Name,
			db.RepositoryColumns.Owner,
			db.RepositoryColumns.CreatedAt,
		),
		db.RepositoryWhere.Owner.EQ(owner),
		db.RepositoryWhere.Name.EQ(name),
	).One(ctx, r.exec)

	if err != nil {
		return nil, err
	}

	return convertRepository(repo), nil
}
