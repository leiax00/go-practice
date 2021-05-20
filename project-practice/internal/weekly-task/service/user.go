package service

import (
	"context"
	v1 "project-practice/api/weekly-task/service/v1"
)

func (service *UserService) GetUserInfo(ctx context.Context, req *v1.UserIdReq) (*v1.UserReply, error) {
	rst := &v1.UserReply{Id: req.Id}
	user, err := service.uc.Get(ctx, req.Id)
	if user != nil {
		rst.Name = user.Name
		rst.Password = user.Password
		rst.Age = int32(user.Age)
		rst.Desc = user.Desc
	}
	return rst, err
}

func (service *UserService) GetUsers(ctx context.Context, req *v1.UserAgeReq) (*v1.UserListReply, error) {
	var rst []*v1.UserReply
	list, err := service.uc.Filter(ctx, req.Age)
	for _, user := range list {
		rst = append(rst, &v1.UserReply{
			Id:       user.Id,
			Name:     user.Name,
			Password: user.Password,
			Age:      int32(user.Age),
			Desc:     user.Desc,
		})
	}
	return &v1.UserListReply{UserReply: rst}, err
}
