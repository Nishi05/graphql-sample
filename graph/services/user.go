package services

import (
	"context"

	"github.com/Nishi05/graphql-sample/graph/db"
	"github.com/Nishi05/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	exec boil.ContextExecutor
}

func convertUser(user *db.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user, err := db.Users( // from users
		qm.Select(db.UserColumns.ID, db.UserColumns.Name), // select id, name
		db.UserWhere.Name.EQ(name),                        // where name = {引数nameの内容}
	).One(ctx, u.exec) // limit 1

	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}
