package main

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"hookfunc/cmd/server/wire"
	"hookfunc/internal/okx"
	"hookfunc/pkg/config"
	"hookfunc/pkg/log"
	"os"
)

// @title						Nunu Example API
// @version					1.0.0
// @description				This is a sample server celler server.
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.url				http://www.swagger.io/support
// @contact.email				support@swagger.io
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:8000
// @securityDefinitions.apiKey	Bearer
// @in							header
// @name						Authorization
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	oneUsdt := 1000000
	// 月： 9U 季： 24U 年： 99U
	subscriptionPrice := []int64{int64(oneUsdt * 9), int64(oneUsdt * 24), int64(oneUsdt * 99)}

	var envConf = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	okxConfig := okx.Config{}
	okxConfig.SubscriptionPrice = subscriptionPrice
	flag.StringVar(&okxConfig.WalletAddress, "WalletAddress", "TS1GYHHFtfP59x6yhb7hizzzATcqkzfsDz", "WalletAddress")
	flag.StringVar(&okxConfig.Server, "ServerUrl", "https://www.okx.com", "ServerUrl")
	flag.IntVar(&okxConfig.Limit, "Limit", 300, "Limit")
	flag.IntVar(&okxConfig.Retry, "Retry", 10, "Retry")
	// 30 = 30USDT
	flag.Parse()

	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewWire(&okxConfig, conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	appConf := os.Getenv("APP_CONF")
	if appConf != "" {
		logger.Info("server use config", zap.String("config name", appConf))
	}

	logger.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))
	logger.Info("docs addr", zap.String("addr", fmt.Sprintf("http://127.0.0.1:%d/swagger/index.html", conf.GetInt("http.port"))))
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
