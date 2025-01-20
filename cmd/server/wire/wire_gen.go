// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"hookfunc/internal/handler"
	"hookfunc/internal/job"
	"hookfunc/internal/okx"
	"hookfunc/internal/repository"
	"hookfunc/internal/server"
	"hookfunc/internal/service"
	"hookfunc/pkg/app"
	"hookfunc/pkg/helper/sid"
	"hookfunc/pkg/jwt"
	"hookfunc/pkg/log"
	"hookfunc/pkg/server/http"
)

// Injectors from wire.go:

func NewWire(config *okx.Config, viperViper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
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
	userInfoService := service.NewUserInfoService(serviceService, repositoryRepository, userInfoRepository)
	userHandler := handler.NewUserHandler(handlerHandler, userInfoService)
	wechatService := service.NewWechatService(serviceService, repositoryRepository, userInfoService)
	wechatHandler := handler.NewWechatHandler(handlerHandler, wechatService)
	strategyRepository := repository.NewStrategyRepository(repositoryRepository)
	strategyService := service.NewStrategyService(serviceService, strategyRepository)
	strategyHandler := handler.NewStrategyHandler(handlerHandler, strategyService)
	barRepository := repository.NewBarRepository(repositoryRepository)
	barService := service.NewBarService(serviceService, barRepository)
	barHandler := handler.NewBarHandler(config, handlerHandler, barService)
	indexRepository := repository.NewIndexRepository(repositoryRepository)
	indexService := service.NewIndexService(config, serviceService, indexRepository)
	indexHandler := handler.NewIndexHandler(handlerHandler, indexService)
	httpServer := server.NewHTTPServer(logger, viperViper, jwtJWT, userHandler, wechatHandler, strategyHandler, barHandler, indexHandler)
	transactionRepository := repository.NewTransactionRepository(repositoryRepository)
	chainService := job.NewChainService(logger, client, transactionRepository, config, userInfoRepository)
	serverJob := server.NewJob(logger, chainService)
	appApp := newApp(httpServer, serverJob)
	return appApp, func() {
	}, nil
}

// wire.go:

var repositorySet = wire.NewSet(repository.NewWechatMiniProgram, repository.NewDB, repository.NewRedis, repository.NewRepository, repository.NewTransaction, repository.NewResourceRepository, repository.NewUserInfoRepository, repository.NewGoodsRepository, repository.NewUserAddressRepository, repository.NewOrderInfoRepository, repository.NewOrderGoodsRepository, repository.NewStrategyRepository, repository.NewBarRepository, repository.NewIndexRepository, repository.NewTransactionRepository)

var serviceSet = wire.NewSet(service.NewService, service.NewResourceService, service.NewUserInfoService, service.NewWechatService, service.NewGoodsService, service.NewUserAddressService, service.NewOrderInfoService, service.NewOrderGoodsService, service.NewStrategyService, service.NewBarService, service.NewIndexService, job.NewChainService)

var handlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler, handler.NewWechatHandler, handler.NewStrategyHandler, handler.NewBarHandler, handler.NewIndexHandler)

var serverSet = wire.NewSet(server.NewHTTPServer, server.NewJob, server.NewTask)

// build App
func newApp(httpServer *http.Server, job2 *server.Job) *app.App {
	return app.NewApp(app.WithServer(httpServer, job2), app.WithName("demo-server"))
}
