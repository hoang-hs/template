package router

import (
	"base/src/common/configs"
	"base/src/present/httpui/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/fx"
)

type RoutersIn struct {
	fx.In
	Engine      *gin.Engine
	AuthService *middlewares.AuthService
}

func RegisterHandler(engine *gin.Engine) {
	// recovery
	engine.Use(middlewares.Recovery())
	//tracer
	engine.Use(otelgin.Middleware(configs.Get().Server.Name))
	engine.Use(middlewares.Tracer())
	// log middleware
	engine.Use(middlewares.Log())
}

func RegisterGinRouters(in RoutersIn) {
	in.Engine.Use(cors.AllowAll())

	group := in.Engine.Group(configs.Get().Server.Http.Prefix)
	group.GET("/ping", middlewares.HealthCheckEndpoint)
	// httpui swagger serve
	if configs.Get().Swagger.Enabled {
		docs.SwaggerInfo.BasePath = fmt.Sprintf("/%s", configs.Get().Server.Http.Prefix)
		group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	registerPublicRouters(group, in)
	protectedGroup := group.Group("/")
	protectedGroup.Use(in.AuthService.AuthenticateServer())
	{
		registerProtectedRouters(protectedGroup, in)
	}
}

func registerPublicRouters(_ *gin.RouterGroup, _ RoutersIn) {

}

func registerProtectedRouters(r *gin.RouterGroup, in RoutersIn) {

}
