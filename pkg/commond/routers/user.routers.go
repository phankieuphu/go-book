package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRouter(rg *gin.RouterGroup) {
	user := rg.Group("/user")
	user.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "user"})
	})
}
