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

func (u ServiceUser) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	panic("implement me")
}

func (u ServiceUser) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	panic("implement me")
}

func (u ServiceUser) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	panic("implement me")
}

func (u ServiceUser) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	panic("implement me")
}

