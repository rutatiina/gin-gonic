package routes

import (
	DemoControllers "wese/core/demo/controllers"
	"wese/core/organization/controllers"
	"wese/core/organization/services"

	"github.com/gin-gonic/gin"
)

func Organization() {
	services.DatabaseConnectAndMigrate()

	r := gin.Default()

	//https://www.stephengream.com/go-nethttp-vs-gin/

	// Routes
	r.GET("/external", DemoControllers.External)
	r.GET("/organizations", controllers.OrganizationList)
	r.GET("/organizations/:id", controllers.OrganizationShow)
	r.POST("/organizations", controllers.OrganizationStore)
	r.PATCH("/organizations/:id", controllers.OrganizationUpdate)
	r.DELETE("/organizations/:id", controllers.OrganizationDestroy)
	r.DELETE("/organizations/:id/void", controllers.OrganizationVoidDestroyed)

	r.GET("/organizations/others", controllers.OrganizationOthers)

	r.Run("localhost:80") //Making sure localhost is there prevents the annoying mac firewall pop up
}
