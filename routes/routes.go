package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/short-circuit-analysis-elements/controllers"
	"github.com/luuisavelino/short-circuit-analysis-elements/middleware"
)

func HandleRequests() {
	router := gin.New()

	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/actuator/health"),
		gin.Recovery(),
		middleware.Logger(),
	)

	actuator := router.Group("/actuator")
	{
		actuator.GET("/health", controllers.HealthGET)
	}

	files := router.Group("/files")
	{
		files.GET("/", controllers.AllFiles)
		files.GET("/:fileId", controllers.OneFile)
	}

	sysInfo := router.Group("/api/files/:fileId")
	{
		sysInfo.GET("/size", controllers.SystemSize)
		sysInfo.GET("/bars", controllers.SystemBars)
		sysInfo.GET("/types", controllers.AllTypes)
		sysInfo.GET("/types/:typeId", controllers.OneType)
	}

	elements := router.Group("/api/files/:fileId/types/:typeId")
	{
		elements.GET("/elements", controllers.AllElementsType)
		elements.GET("/elements/:elementId/", controllers.OneElement)
	}

	router.Run(":8080")
}