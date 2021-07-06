package interfaces

import (
	"golang-k8s-helm-helloworld/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
}

func MessageGetHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		h := dto.Message{}
		h.Text = "Hello world"
		c.JSON(http.StatusOK, h)
	}
}
