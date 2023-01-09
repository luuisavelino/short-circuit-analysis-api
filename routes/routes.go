package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/short-circuit-analysis-elements/controllers"
	"github.com/luuisavelino/short-circuit-analysis-elements/middleware"
)

func HandleRequests() {
	r := gin.Default()
	r.Use(middleware.Logger())

	r.GET("/health/liveness", controllers.Liveness)
	r.GET("/health/readiness", controllers.Readiness)
	r.GET("/api/files", controllers.AllFiles)
	r.GET("/api/files/:fileId", controllers.OneFile)
	r.GET("/api/files/:fileId/size", controllers.SystemSize)
	r.GET("/api/files/:fileId/bars", controllers.SystemBars)
	r.GET("/api/files/:fileId/elements", controllers.AllElements)
	r.GET("/api/files/:fileId/elements/type/:typeId", controllers.AllElementsType)
	r.GET("/api/files/:fileId/elements/type/:typeId/element/:elementId", controllers.OneElement)

	r.Run(":8080")
}
