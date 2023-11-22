// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"hookfunc-media/internal/repository"
	"hookfunc-media/internal/server"
	"hookfunc-media/pkg/app"
	"hookfunc-media/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func NewWire(viperViper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	db := repository.NewDB(viperViper, logger)
	migrate := server.NewMigrate(db, logger)
	appApp := newApp(migrate)
	return appApp, func() {
	}, nil
}

// wire.go:

var repositorySet = wire.NewSet(repository.NewDB, repository.NewRedis, repository.NewRepository, repository.NewUserRepository)

// build App
func newApp(migrate *server.Migrate) *app.App {
	return app.NewApp(app.WithServer(migrate), app.WithName("demo-migrate"))
}
