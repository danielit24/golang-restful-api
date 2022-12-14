package app

import (
	"database/sql"
	"time"

	"github.com/danielit24/golang-restful-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "daniel:password@tcp(localhost:3306)/golang_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
