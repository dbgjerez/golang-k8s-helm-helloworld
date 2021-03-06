package interfaces

import (
	"golang-k8s-helm-helloworld/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthcheckHandler struct {
}

func HealthcheckGetHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		h := dto.Health{}
		h.Status = dto.HealhStatusUp
		c.JSON(http.StatusOK, h)
	}
}
