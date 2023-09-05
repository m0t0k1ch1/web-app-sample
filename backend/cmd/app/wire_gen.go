// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"app/db"
	"app/env"
	"app/service"
	"app/service/app/v1"
	"context"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeApp(ctx context.Context, conf Config) (*App, error) {
	server := conf.Server
	mySQL := conf.MySQL
	connection, err := db.NewConnection(mySQL)
	if err != nil {
		return nil, err
	}
	container := env.NewContainer(connection)
	base := service.NewBase(container)
	taskService := appv1.NewTaskService(base)
	mainServer := NewServer(server, taskService)
	app := NewApp(mainServer)
	return app, nil
}
