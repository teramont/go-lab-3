package main

import (
	"github.com/teramont/go-lab-3/server/storage"
)

func Compose(port HttpPortNumber) (*ApiServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	store := storage.NewStorage(db)
	server := NewServer(port, storage.HttpHandler(store))
	return &server, nil
}
