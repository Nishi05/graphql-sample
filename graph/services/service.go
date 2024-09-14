package services

import (
	"context"

	"github.com/Nishi05/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	ListUsersByID(ctx context.Context, IDs []string) ([]*model.User, error)
}

type RepoService interface {
	GetRepoByID(ctx context.Context, id string) (*model.Repository, error)
	GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error)
}

type IssueService interface {
	GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
	ListIssueInRepository(ctx context.Context, repoID string, after *string, before *string, first *int, last *int) (*model.IssueConnection, error)
}

type PullRequestService interface {
	GetPullRequestByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.PullRequest, error)
}

type Services interface {
	UserService
	RepoService
	IssueService
	PullRequestService
}

type services struct {
	*userService
	*repoService
	*issueService
	*pullRequestService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		&userService{
			exec: exec,
		},
		&repoService{
			exec: exec,
		},
		&issueService{
			exec: exec,
		},
		&pullRequestService{
			exec: exec,
		},
	}
}
