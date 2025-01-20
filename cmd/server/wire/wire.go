//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"hookfunc/internal/handler"
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

var repositorySet = wire.NewSet(
	repository.NewWechatMiniProgram,
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewResourceRepository,
	repository.NewUserInfoRepository,
	repository.NewGoodsRepository,
	repository.NewUserAddressRepository,
	repository.NewOrderInfoRepository,
	repository.NewOrderGoodsRepository,
	repository.NewStrategyRepository,
	repository.NewBarRepository,
	repository.NewIndexRepository,
	repository.NewTransactionRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewResourceService,
	service.NewUserInfoService,
	service.NewWechatService,
	service.NewGoodsService,
	service.NewUserAddressService,
	service.NewOrderInfoService,
	service.NewOrderGoodsService,
	service.NewStrategyService,
	service.NewBarService,
	service.NewIndexService,
	service.NewChainService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewWechatHandler,
	handler.NewStrategyHandler,
	handler.NewBarHandler,
	handler.NewIndexHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
	server.NewTask,
)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*okx.Config, *viper.Viper, *log.Logger) (*app.App, func(), error) {

	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
