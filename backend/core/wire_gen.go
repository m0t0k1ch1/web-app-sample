// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package core

import (
	"context"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeApp(ctx context.Context, confPath ConfigPath) (*App, error) {
	appConfig, err := provideAppConfig(confPath)
	if err != nil {
		return nil, err
	}
	handler, err := provideSentryHandler(appConfig)
	if err != nil {
		return nil, err
	}
	clock := provideClock()
	mySQLContainer, err := provideMySQLContainer(appConfig)
	if err != nil {
		return nil, err
	}
	taskService := provideTaskService(clock, mySQLContainer)
	server := provideServer(appConfig, handler, taskService)
	app := provideApp(appConfig, server)
	return app, nil
}
