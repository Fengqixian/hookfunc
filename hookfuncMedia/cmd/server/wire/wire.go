//go:build wireinject
// +build wireinject

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

var repositorySet = wire.NewSet(
	repository.NewWechatMiniProgram,
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewResourceRepository,
	repository.NewUserInfoRepository,
	repository.NewGoodsRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewResourceService,
	service.NewUserInfoService,
	service.NewWechatService,
	service.NewGoodsService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewResourceHandler,
	handler.NewWechatHandler,
	handler.NewGoodsHandler,
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
