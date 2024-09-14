//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"hookfunc/internal/handler"
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
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewResourceHandler,
	handler.NewWechatHandler,
	handler.NewGoodsHandler,
	handler.NewUserAddressHandler,
	handler.NewOrderInfoHandler,
	handler.NewOrderGoodsHandler,
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

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {

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
