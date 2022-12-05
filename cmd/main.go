package main

import (
	"github.com/gin-gonic/gin"
	"github.com/patrick-blip/vocher-init/pkg/commond/config"
	"github.com/patrick-blip/vocher-init/pkg/commond/controllers"
	"github.com/patrick-blip/vocher-init/pkg/commond/models"
	"github.com/patrick-blip/vocher-init/pkg/commond/routers"
)

var r = gin.Default()

func main() {
	config.LoadConfig("../pkg/commond/envs/.env")
	r.GET("/", controllers.FindBook)
	r.POST("/", controllers.CreateBook)
	r.PATCH("/:id", controllers.UpdateBook)
	r.DELETE("/:id", controllers.DeleteBook)
	models.ConnectDataBase()
	getRouters()
	// port := viper.Get("PORT").(string)
	r.Run("0.0.0.0:4000")
}

func getRouters() {
	v1 := r.Group("/v1")
	routers.UserRouter(v1)
	routers.VocherRouter(v1)
}
