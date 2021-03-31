package models

import "time"

type Activities struct {
	Id           int
	UserId       int
	ActivityType string
	TimeSpent    string
	CreatedAt    time.Time
}

type ActivityService interface {
	Play()
	Sleep()
	Eat()
	Read()
}

type ActivityRepository interface {
	GetActivityId(id int) (*Activities, error)
	GetActivityByUser(userId int) ([]Activities, error)
	CreateActivity(act Activities) (int, error)
	UpdateActivity(act Activities) (int, error)
	DeleteActivity(id int) (int, error)
}

func (a Activities) IsValid() error {
	return nil
}

func (a Activities) IsDone() bool {
	return false
}
