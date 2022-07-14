package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsHandlerGetHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		prometheusHandler := promhttp.Handler()
		prometheusHandler.ServeHTTP(c.Writer, c.Request)
	}
}
