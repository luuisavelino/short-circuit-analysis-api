package controllers

import "github.com/gin-gonic/gin"

func Readiness(c *gin.Context) {
	// Readiness probe do kubernetes

}

func Liveness(c *gin.Context) {
	// Liveness probe do kubernetes
}