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

	files := router.Group("/api/files")
	{
		files.GET("/", controllers.AllFiles)
		files.GET("/:fileId", controllers.OneFile)
	}

	system := router.Group("/api/files/:fileId")
	{
		system.GET("/size", controllers.SystemSize)
		system.GET("/bars", controllers.SystemBars)
		system.GET("/elements", controllers.AllElements)
		system.GET("/elements/type/:typeId", controllers.AllElementsType)
		system.GET("/elements/type/:typeId/element/:elementId", controllers.OneElement)
	}

	router.Run(":8080")
}
