package main

import (
	"golang-k8s-helm-helloworld/interfaces"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", interfaces.HealthcheckGetHandler())
		v1.GET("/greetings", interfaces.MessageGetHandler())
	}

	router.GET("/metrics", interfaces.MetricsHandlerGetHandler())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "Not found"})
	})

	router.Run(":8080")
}
