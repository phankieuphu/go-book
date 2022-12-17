package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/patrick-blip/vocher-init/pkg/commond/controllers"
)

func TimeRouter(rg *gin.RouterGroup) {
	time := rg.Group("/time")
	time.GET("/", controllers.CompareTime)
}
