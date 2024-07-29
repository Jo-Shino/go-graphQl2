package services

import (
	"context"

	"github.com/saki-engineering/graphql-sample/graph/model"
	"github.com/volatiletech/sqlboiler/boil"
)

type Services interface {
	UserService
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type services struct {
	*userService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:        &userService{exec: exec},
	}
}