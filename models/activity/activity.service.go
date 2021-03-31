package activity

import (
	"activity/models"
	"activity/pb"
	"context"
	"time"
)

type Service struct {
	repo   models.ActivityRepository
}

func NewService(repo models.ActivityRepository) models.ActivityService {
	return &Service{repo: repo}
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

func (s Service) CreateActivity(ctx context.Context, request *pb.ActivityRequestBdy) (*pb.Response, error) {
	res,err := s.repo.CreateActivity(mapActivityRequestToActivityModel(request))
	if err != nil {
		return nil, err
	}
	return &pb.Response{Id: int64(res)},nil
}

func (s Service) GetActivity(ctx context.Context, request *pb.GetActivityRequest) (*pb.GetActivityResponse, error) {
	res,err := s.repo.GetActivityId(int(request.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetActivityResponse{Activity: mapActivityModelToActivityResponse(*res)},nil
}

func (s Service) GetActivityByUserId(ctx context.Context, request *pb.GetActivityRequest) (*pb.GetActivityByUserResponse, error) {
	res,err := s.repo.GetActivityByUser(int(request.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetActivityByUserResponse{Activity: mapActivityModelsToActivitiesResponse(res)},nil
}

func (s Service) UpdateActivity(ctx context.Context, request *pb.ActivityRequestBdy) (*pb.Response, error) {
	res,err := s.repo.UpdateActivity(mapActivityRequestToActivityModel(request))
	if err != nil {
		return nil, err
	}
	return &pb.Response{Id: int64(res)},nil
}

func (s Service) DeleteActivity(ctx context.Context, request *pb.DeleteActivityRequest) (*pb.Response, error) {
	res,err := s.repo.DeleteActivity(int(request.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Response{Id: int64(res)},nil
}


func mapActivityRequestToActivityModel(requestBdy *pb.ActivityRequestBdy) models.Activities{
	return models.Activities{
		UserId:       int(requestBdy.Request.UserId),
		ActivityType: requestBdy.Request.ActivityType,
		TimeSpent:    requestBdy.Request.TimeSpent,
		CreatedAt:    time.Time{},
	}
}

func mapActivityModelToActivityResponse(activity models.Activities) *pb.ActivityResponse{
	return &pb.ActivityResponse{
		Id:           int64(activity.Id),
		UserId:       int64(activity.UserId),
		ActivityType: activity.ActivityType,
		TimeSpent:    activity.TimeSpent,
		CreatedAt:    activity.CreatedAt.String(),
	}
}

func mapActivityModelsToActivitiesResponse(activity []models.Activities) []*pb.ActivityResponse{
	var response []*pb.ActivityResponse

	for _,value := range activity {
		response = append(response,&pb.ActivityResponse{
			Id:           int64(value.Id),
			UserId:       int64(value.UserId),
			ActivityType: value.ActivityType,
			TimeSpent:    value.TimeSpent,
			CreatedAt:    value.CreatedAt.String(),
		})
	}

	return response
}