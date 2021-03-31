package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDBConnection(dns string) (*sql.DB,error) {
	db, err := sql.Open("mysql", dns)

	if err != nil {
		return nil, err
	}
    err = db.Ping()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return db,nil
}