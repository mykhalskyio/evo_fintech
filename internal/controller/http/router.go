package http

import (
	_ "evo_fintech/docs"
	"evo_fintech/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(service *service.DataService) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.MaxMultipartMemory = 5368709120
	dataController := newDataController(service)
	api := router.Group("/api")
	{
		api.POST("/upload", dataController.upload)
		api.GET("/download/:format", dataController.download)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
