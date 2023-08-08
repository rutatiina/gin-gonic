package routes

import (
	DemoControllers "wese/core/demo/controllers"
	"wese/core/module_template/controllers"
	"wese/core/module_template/services"

	"github.com/gin-gonic/gin"
)

func RouteName() {
	services.DatabaseConnectAndMigrate()

	r := gin.Default()

	//https://www.stephengream.com/go-nethttp-vs-gin/

	// Routes
	r.GET("/external", DemoControllers.External)
	r.GET("/resource-name", controllers.ControllerNameList)
	r.GET("/resource-name/:id", controllers.ControllerNameShow)
	r.POST("/resource-name", controllers.ControllerNameStore)
	r.PATCH("/resource-name/:id", controllers.ControllerNameUpdate)
	r.DELETE("/resource-name/:id", controllers.ControllerNameDestroy)
	r.DELETE("/resource-name/:id/void", controllers.ControllerNameVoidDestroyed)

	r.GET("/resource-name/others", controllers.ControllerNameOthers)

	r.Run("localhost:80") //Making sure localhost is there prevents the annoying mac firewall pop up
}
