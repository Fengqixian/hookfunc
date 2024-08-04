// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"hookfunc-media/internal/handler"
	"hookfunc-media/internal/repository"
	"hookfunc-media/internal/server"
	"hookfunc-media/internal/service"
	"hookfunc-media/pkg/app"
	"hookfunc-media/pkg/helper/sid"
	"hookfunc-media/pkg/jwt"
	"hookfunc-media/pkg/log"
	"hookfunc-media/pkg/server/http"
)

// Injectors from wire.go:

func NewWire(viperViper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	jwtJWT := jwt.NewJwt(viperViper)
	handlerHandler := handler.NewHandler(logger)
	miniProgram := repository.NewWechatMiniProgram(viperViper)
	db := repository.NewDB(viperViper, logger)
	client := repository.NewRedis(viperViper)
	repositoryRepository := repository.NewRepository(miniProgram, db, client, logger)
	transaction := repository.NewTransaction(repositoryRepository)
	sidSid := sid.NewSid()
	serviceService := service.NewService(transaction, logger, sidSid, jwtJWT)
	userInfoRepository := repository.NewUserInfoRepository(repositoryRepository)
	userInfoService := service.NewUserInfoService(serviceService, userInfoRepository)
	userHandler := handler.NewUserHandler(handlerHandler, userInfoService)
	wechatService := service.NewWechatService(serviceService, repositoryRepository, userInfoService)
	wechatHandler := handler.NewWechatHandler(handlerHandler, wechatService)
	httpServer := server.NewHTTPServer(logger, viperViper, jwtJWT, userHandler, wechatHandler)
	job := server.NewJob(logger)
	appApp := newApp(httpServer, job)
	return appApp, func() {
	}, nil
}

// wire.go:

var repositorySet = wire.NewSet(repository.NewWechatMiniProgram, repository.NewDB, repository.NewRedis, repository.NewRepository, repository.NewTransaction, repository.NewResourceRepository, repository.NewUserInfoRepository)

var serviceSet = wire.NewSet(service.NewService, service.NewResourceService, service.NewUserInfoService, service.NewWechatService)

var handlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler, handler.NewResourceHandler, handler.NewWechatHandler)

var serverSet = wire.NewSet(server.NewHTTPServer, server.NewJob, server.NewTask)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(app.WithServer(httpServer, job), app.WithName("demo-server"))
}
