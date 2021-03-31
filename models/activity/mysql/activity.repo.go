package mysql

import (
	"aishW/models"
	"database/sql"
	"time"
)

const (
	insertActivity       = `INSERT INTO ACTIVITY (USER_ID,ACTIVITY_TYPE,DURATION,CREATED_AT) VALUES (?,?,?)`
	selectActivity       = `SELECT * FROM ACTIVITY WHERE ID = ?`
	selectActivityByUser = `SELECT * FROM ACTIVITY WHERE USER_ID = ?`
	updateActivity       = `UPDATE ACTIVITY SET USER_ID=?,ACTIVITY_TYPE=?,DURATION=?,UPDATED_AT WHERE ID=?`
	deleteActivity       = `DELETE ACTIVITY WHERE ID=?`
)

type activityRepo struct {
	db *sql.DB
}

func NewActivityRepo(db *sql.DB) models.ActivityRepository {
	return activityRepo{db: db}
}

func (a activityRepo) GetActivityId(id int) (*models.Activities, error) {
	var act models.Activities
	rows, err := a.db.Query(selectActivity, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		scanError := rows.Scan(&act)
		if scanError != nil {
			return nil, scanError
		}
	}
	return &act, nil
}

func (a activityRepo) GetActivityByUser(userId int) ([]models.Activities, error) {
	var activities []models.Activities
	rows, err := a.db.Query(selectActivityByUser, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var act models.Activities
		scanError := rows.Scan(&act)
		if scanError != nil {
			return nil, scanError
		}
		activities = append(activities, act)
	}
	return activities, nil
}

func (a activityRepo) CreateActivity(act models.Activities) (int, error) {
	res, err := a.db.Exec(insertActivity, act.UserId, act.ActivityType, act.TimeSpent, time.Now())
	if err != nil {
		return 0, err
	}
	if affected, err := res.RowsAffected(); affected <= 0 || err != nil {
		return 0, nil
	}
	id, _ := res.LastInsertId()
	return int(id), nil
}

func (a activityRepo) UpdateActivity(act models.Activities) (int, error) {
	res, err := a.db.Exec(updateActivity, act.UserId, act.ActivityType, act.TimeSpent, time.Now())
	if err != nil {
		return 0, err
	}
	if affected, err := res.RowsAffected(); affected <= 0 || err != nil {
		return 0, nil
	}
	id, _ := res.LastInsertId()
	return int(id), nil
}

func (a activityRepo) DeleteActivity(id int) (int, error) {
	res, err := a.db.Exec(deleteActivity, id)
	if err != nil {
		return 0, err
	}
	if affected, err := res.RowsAffected(); affected <= 0 || err != nil {
		return 0, nil
	}
	affectedId, _ := res.LastInsertId()
	return int(affectedId), nil
}
