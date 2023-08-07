package routes

import (
	"wese/_module_template_/controllers"
	"wese/_module_template_/services"

	"github.com/gin-gonic/gin"
)

func _router_name_() {
	services.DatabaseConnectAndMigrate()

	r := gin.Default()

	//https://www.stephengream.com/go-nethttp-vs-gin/

	// Routes
	r.GET("/users", controllers.List)
	r.GET("/users/:id", controllers.Show)
	r.POST("/users", controllers.Store)
	r.PATCH("/users/:id", controllers.Update)
	r.DELETE("/users/:id", controllers.Destroy)
	r.DELETE("/users/:id/void", controllers.VoidDestroyed)

	r.GET("/others", controllers.Others)

	r.Run("localhost:8081") //Making sure localhost is there prevents the annoying mac firewall pop up
}
