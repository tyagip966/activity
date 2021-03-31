package mysql

import (
	"aishW/models"
	"database/sql"
)

type userRepo struct {
	db   *sql.DB
}

func NewUserRepo(db *sql.DB) models.UserRepository {
	return userRepo{db: db}
}