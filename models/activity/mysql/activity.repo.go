package mysql

import (
	"aishW/models"
	"database/sql"
)

type activityRepo struct {
	db   *sql.DB
}

func NewActivityRepo(db *sql.DB) models.ActivityRepository {
	return activityRepo{db: db}
}