package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"project-practice/internal/weekly-task/biz"
	"project-practice/internal/weekly-task/data/ent/user"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper("data/user", logger),
	}
}

func (repo *userRepo) GetUserInfo(ctx context.Context, id int64) (*biz.User, error) {
	tmp, err := repo.data.db.User.Get(ctx, id)
	if tmp == nil {
		return &biz.User{
			Id: id,
		}, err
	}
	return &biz.User{
		Id:       tmp.ID,
		Name:     tmp.Username,
		Password: tmp.Password,
		Age:      int8(tmp.Age),
		Desc:     tmp.Desc,
	}, err
}

func (repo *userRepo) GetUsers(ctx context.Context, age int32) ([]*biz.User, error) {
	var rst []*biz.User
	users, err := repo.data.db.User.Query().Where(user.AgeEQ(age)).All(ctx)
	for _, u := range users {
		rst = append(rst, &biz.User{
			Id:       u.ID,
			Name:     u.Username,
			Password: u.Password,
			Age:      int8(u.Age),
			Desc:     u.Desc,
		})
	}
	return rst, err
}
