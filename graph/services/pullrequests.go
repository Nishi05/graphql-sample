package services

import (
	"context"
	"log"

	"github.com/Nishi05/graphql-sample/graph/db"
	"github.com/Nishi05/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type pullRequestService struct {
	exec boil.ContextExecutor
}

func convertPullRequest(pullRequest *db.Pullrequest) *model.PullRequest {
	prURL, err := model.UnmarshalURI(pullRequest.URL)
	if err != nil {
		log.Println("invalid URL", pullRequest.URL)
	}
	return &model.PullRequest{
		ID:          pullRequest.ID,
		BaseRefName: pullRequest.BaseRefName,
		Closed:      (pullRequest.Closed == 1),
		HeadRefName: pullRequest.HeadRefName,
		URL:         prURL,
		Number:      int(pullRequest.Number),
		Repository:  &model.Repository{ID: pullRequest.Repository},
	}
}

func (p *pullRequestService) GetPullRequestByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.PullRequest, error) {
	pullRequest, err := db.Pullrequests(
		qm.Select(
			db.PullrequestColumns.ID,
			db.PullrequestColumns.BaseRefName,
			db.PullrequestColumns.Closed,
			db.PullrequestColumns.HeadRefName,
			db.PullrequestColumns.Number,
			db.PullrequestColumns.Repository,
		),
		db.PullrequestWhere.Repository.EQ(repoID),
		db.PullrequestWhere.Number.EQ(int64(number)),
	).One(ctx, p.exec)

	if err != nil {
		return nil, err
	}
	return convertPullRequest(pullRequest), nil

}
