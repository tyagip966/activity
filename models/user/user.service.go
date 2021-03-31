package user

import (
	"aishW/models"
	"aishW/pb"
	"context"
)

type ServiceUser struct {
	repo models.UserRepository
}

func NewServiceUser(repo models.UserRepository) *ServiceUser {
	return &ServiceUser{repo: repo}
}

func (u ServiceUser) CreateUser(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	res,err := u.repo.CreateUser(mapUserRequestToUserModel(request))
	if err != nil {
		return nil, err
	}
	return &pb.Response{Id: int64(res)},nil
}

func (u ServiceUser) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res,err := u.repo.GetUser(int(request.Userid))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{User: mapUserModelToUserResponse(*res)},nil
}

func (u ServiceUser) UpdateUser(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	res,err := u.repo.UpdateUser(mapUserRequestToUserModel(request))
	if err != nil {
		return nil, err
	}
	return &pb.Response{Id: int64(res)},nil
}

func (u ServiceUser) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.Response, error) {
	res,err := u.repo.DeleteUser(int(request.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Response{Id: int64(res)},nil
}

func mapUserRequestToUserModel(user *pb.Request) models.User{
	return models.User{
		Name:  user.Request.Name,
		Email: user.Request.Email,
		Phone: user.Request.Phone,
	}
}

func mapUserModelToUserResponse(user models.User) *pb.UserResponse{
	return &pb.UserResponse{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}
}