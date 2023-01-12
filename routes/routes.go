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

	files := router.Group("/api/v2/files")
	{
		files.GET("/", controllers.AllFiles)
		files.GET("/:fileId", controllers.OneFile)
	}

	sysInfo := files.Group("/:fileId")
	{
		sysInfo.GET("/size", controllers.SystemSize)
		sysInfo.GET("/bars", controllers.SystemBars)
		sysInfo.GET("/types", controllers.AllTypes)
		sysInfo.GET("/types/:typeId", controllers.OneType)
	}

	elements := sysInfo.Group("/types/:typeId")
	{
		elements.GET("/elements", controllers.AllElementsType)
		elements.GET("/elements/:elementId", controllers.OneElement)
	}

	router.Run(":8080")
}