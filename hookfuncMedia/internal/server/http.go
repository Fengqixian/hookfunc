package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	apiV1 "hookfunc-media/api/v1"
	"hookfunc-media/docs"
	"hookfunc-media/internal/handler"
	"hookfunc-media/internal/middleware"
	"hookfunc-media/pkg/jwt"
	"hookfunc-media/pkg/log"
	"hookfunc-media/pkg/server/http"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	wechatHandler *handler.WechatHandler,
	goodsHandler *handler.GoodsHandler,
	userAddressHandler *handler.UserAddressHandler,
	orderInfoHandler *handler.OrderInfoHandler,
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
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
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
			":)": "Thank you for using nunu!",
		})
	})

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/wechat/qr/login", wechatHandler.ProgramQrCodeLogin)
			noAuthRouter.POST("/wechat/program/login", wechatHandler.ProgramLogin)
		}
		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		{

			noStrictAuthRouter.POST("/goods/info", goodsHandler.Info)
			noStrictAuthRouter.GET("/goods/list", goodsHandler.List)

			noStrictAuthRouter.POST("/order/place", orderInfoHandler.Place)
			noStrictAuthRouter.POST("/order/cancel", orderInfoHandler.Cancel)
			noStrictAuthRouter.GET("/order/list", orderInfoHandler.List)

			noStrictAuthRouter.GET("/user/address/list", userAddressHandler.List)
			noStrictAuthRouter.POST("/user/address/update", userAddressHandler.Update)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.GET("/user", userHandler.GetProfile)
			strictAuthRouter.POST("/user/address/create", userAddressHandler.Create)

		}
	}

	return s
}
