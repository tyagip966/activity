package activity

import (
	"aishW/models"
	"aishW/pb"
	"context"
)

type Service struct {
}

func NewService() models.ActivityService {
	return &Service{}
}

func (s Service) Play() {
	panic("implement me")
}

func (s Service) Sleep() {
	panic("implement me")
}

func (s Service) Eat() {
	panic("implement me")
}

func (s Service) Read() {
	panic("implement me")
}

func (s Service) CreateActivity(ctx context.Context, request *pb.CreateActivityRequest) (*pb.CreateActivityResponse, error) {
	panic("implement me")
}

func (s Service) GetActivity(ctx context.Context, request *pb.GetActivityRequest) (*pb.GetActivityResponse, error) {
	panic("implement me")
}

func (s Service) UpdateActivity(ctx context.Context, request *pb.UpdateActivityRequest) (*pb.UpdateActivityResponse, error) {
	panic("implement me")
}

func (s Service) DeleteActivity(ctx context.Context, request *pb.DeleteActivityRequest) (*pb.DeleteActivityResponse, error) {
	panic("implement me")
}
