package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	apiV1 "hookfunc/api/v1"
	"hookfunc/docs"
	"hookfunc/internal/handler"
	"hookfunc/internal/middleware"
	"hookfunc/pkg/jwt"
	"hookfunc/pkg/log"
	"hookfunc/pkg/server/http"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	wechatHandler *handler.WechatHandler,
	strategyHandler *handler.StrategyHandler,
	barHandler *handler.BarHandler,
	indexHandler *handler.IndexHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "rain or shine",
		})
	})

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/wechat/qr/login", wechatHandler.ProgramQrCodeLogin)
			noAuthRouter.POST("/wechat/program/login", wechatHandler.ProgramLogin)
			noAuthRouter.POST("/sms/code", userHandler.SendSmsCode)
			noAuthRouter.POST("/verification/sms/code", userHandler.VerificationSmsCode)

		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.GET("/user", userHandler.GetProfile)

		}

		// Strict permission routing group
		strategyRouter := v1.Group("/strategy").Use(middleware.StrictAuth(jwt, logger))
		{
			strategyRouter.GET("/list", strategyHandler.ListStrategy)
			strategyRouter.POST("/index/list", strategyHandler.ListStrategyIndex)
			strategyRouter.POST("/index/delete", strategyHandler.DeleteStrategyIndex)
			strategyRouter.POST("/delete", strategyHandler.DeleteStrategy)
			strategyRouter.POST("/create", strategyHandler.CreateStrategy)
			strategyRouter.POST("/subscription", strategyHandler.UpdateStrategySubscriptionState)
		}

		// Strict permission routing group
		indexRouter := v1.Group("/index")
		{
			noAuthRouter.GET("/bar/list", barHandler.ListBar)
			indexRouter.GET("/list", indexHandler.ListIndex)

		}

	}

	return s
}
