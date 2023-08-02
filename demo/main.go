package main

import (
	"wese/demo/controllers"
	"wese/demo/services"

	"github.com/gin-gonic/gin"
)

func main() {
	services.DatabaseConnectAndMigrate()

	r := gin.Default()

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
