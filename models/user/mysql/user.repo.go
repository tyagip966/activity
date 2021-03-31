package mysql

import (
	"aishW/models"
	"database/sql"
)

const (
	insertUser = `INSERT INTO USER (NAME,EMAIL,PHONE) VALUES (?,?,?)`
	selectUser = `SELECT * FROM USER WHERE ID = ?`
	updateUser = `UPDATE USER SET NAME=?,EMAIL=?,PHONE=? WHERE ID=?`
	deleteUser = `DELETE USER WHERE ID=?`
)

type userRepo struct {
	db   *sql.DB
}

func NewUserRepo(db *sql.DB) models.UserRepository {
	return userRepo{db: db}
}

func (u userRepo) GetUser(id int) (*models.User, error) {
	var user models.User
	rows,err := u.db.Query(selectUser,id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		scanError := rows.Scan(&user)
		if scanError != nil {
			return nil, scanError
		}
	}
	return &user, nil
}

func (u userRepo) CreateUser(user models.User) (int, error) {
	res,err := u.db.Exec(insertUser,user.Name,user.Email,user.Phone)
	if err != nil {
		return 0, err
	}
	if affected,err := res.RowsAffected(); affected <= 0 || err != nil {
		return 0,nil
	}
	id,_ := res.LastInsertId()
	return int(id), nil
}

func (u userRepo) UpdateUser(user models.User) (int, error) {
	res,err := u.db.Exec(updateUser,user.Name,user.Email,user.Phone)
	if err != nil {
		return 0, err
	}
	if affected,err := res.RowsAffected(); affected <= 0 || err != nil {
		return 0,nil
	}
	id,_ := res.LastInsertId()
	return int(id), nil
}

func (u userRepo) DeleteUser(id int) (int, error) {
	res,err := u.db.Exec(deleteUser,id)
	if err != nil {
		return 0, err
	}
	if affected,err := res.RowsAffected(); affected <= 0 || err != nil {
		return 0,nil
	}
	affectedId,_ := res.LastInsertId()
	return int(affectedId), nil
}