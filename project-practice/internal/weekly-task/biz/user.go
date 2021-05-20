package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int64
	Name     string
	Password string
	Age      int8
	Desc     string
}

type UserRepo interface {
	GetUserInfo(ctx context.Context, id int64) (*User, error)
	GetUsers(ctx context.Context, age int32) ([]*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper("useCase/user", logger)}
}

func (useCase *UserUseCase) Get(ctx context.Context, id int64) (*User, error) {
	return useCase.repo.GetUserInfo(ctx, id)
}

func (useCase *UserUseCase) Filter(ctx context.Context, age int32) ([]*User, error) {
	return useCase.repo.GetUsers(ctx, age)
}
