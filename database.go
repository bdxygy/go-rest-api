package main

import (
	"database/sql"
	"github.com/bdxygy/exception"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golangapi")
	exception.Throw(err)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	return db
}
