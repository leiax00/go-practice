package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"project-practice/api/weekly-task/service/v1"
	"project-practice/internal/weekly-task/biz"
)

var ProviderSet = wire.NewSet(NewCartService)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewCartService(useCase *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  useCase,
		log: log.NewHelper("service/cart", logger)}
}
