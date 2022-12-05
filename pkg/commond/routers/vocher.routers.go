package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VocherRouter(rg *gin.RouterGroup) {
	vocher := rg.Group("/vocher")
	vocher.GET("/", func(c *gin.Context) {
		fullLink := c.FullPath()
		c.JSON(http.StatusOK, gin.H{"data": fullLink})
	})
}
