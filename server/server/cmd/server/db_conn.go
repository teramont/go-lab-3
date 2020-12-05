package main

import (
	"database/sql"

	"github.com/teramont/go-lab-3/server/db"
)

func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     "lab3",
		User:       "lab3",
		Password:   "lab3",
		Host:       "localhost",
		DisableSSL: true,
	}
	return conn.Open()
}
