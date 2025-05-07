package api

import (
	"update-api/docs"
	"update-api/pkg/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerAPI struct {
	eng *gin.Engine

	serverHandler ServerHandler
	resultHandler ResultHandler
}

func NewServerAPI(serviceCollector services.ServiceCollector) *ServerAPI {
	gin.SetMode(gin.DebugMode)
	return &ServerAPI{
		eng:           gin.New(),
		serverHandler: NewDefaultServerHandler(serviceCollector.ServerService()),
		resultHandler: NewDefaultResultHandler(serviceCollector.ResultService()),
	}
}

func (inst *ServerAPI) Start(address string) error {
	// swagger info setting
	docs.SwaggerInfo.Title = "Update API Swagger"
	docs.SwaggerInfo.Description = "This is a backend for managing a server list and settings of update-service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = address
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	inst.eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// server handlers
	apiGroup := inst.eng.Group("/api/v1")
	apiGroup.GET("/server/:uuid", inst.serverHandler.Get)
	apiGroup.GET("servers", inst.serverHandler.GetAll)
	apiGroup.DELETE("/server/:uuid", inst.serverHandler.Delete)
	apiGroup.PATCH("/server/:uuid", inst.serverHandler.Patch)
	apiGroup.POST("/server", inst.serverHandler.Post)

	// result handlers
	apiGroup.GET("/result/:uuid", inst.resultHandler.Get)
	apiGroup.GET("/results", inst.resultHandler.GetAll)

	return inst.eng.Run(address)
}
