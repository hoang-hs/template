package bootstrap

import (
	"base/src/common/configs"
	"base/src/present/httpui/router"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func BuildHTTPServerModule() fx.Option {
	return fx.Options(
		fx.Provide(gin.New),
		fx.Invoke(router.RegisterHandler),
		fx.Invoke(router.RegisterGinRouters),
		fx.Invoke(NewHttpServer),
	)
}

func NewHttpServer(logger *zap.SugaredLogger, lc fx.Lifecycle, engine *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := engine.Run(fmt.Sprintf(":%s", configs.Get().Server.Http.Address)); err != nil {
					logger.Fatalf("Cannot start application due by error [%v]", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Infof("Stopping HTTP server")
			return nil
		},
	})
}
