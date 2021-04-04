package models

import (
	"errors"
	"log"
	"strings"
	"time"
)

type Activities struct {
	Id           int
	UserId       int
	ActivityType string
	TimeSpent    string
	Status       string
	Label        string
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

func (a *Activities) IsValid() error {
	log.Print("jkxcvxjvx")
	if strings.ToLower(a.ActivityType) == "sleep" && (a.TimeSpent < "6h" || a.TimeSpent > "8h") {
         return errors.New("invalid time spent for activity "+a.ActivityType)
	}
	if strings.ToLower(a.ActivityType) == "play" && (a.TimeSpent < "2h" || a.TimeSpent > "2h") {
		return errors.New("invalid time spent for activity "+a.ActivityType)
	}
	if strings.ToLower(a.ActivityType) == "eat" && (a.TimeSpent < "1h" || a.TimeSpent > "1h") {
		return errors.New("invalid time spent for activity "+a.ActivityType)
	}
	return nil
}

func (a *Activities) IsDone() bool {
	 log.Println("Hi Golang")
     if strings.ToLower(a.Status) == "done" {
     	return true
	 }
	return false
}
