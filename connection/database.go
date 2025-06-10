package connection

import (
	"database/sql"
	"go-restful-api/helper"
	"time"
)

func DatabaseConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
