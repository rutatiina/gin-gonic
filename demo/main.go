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

	/*
		r.GET("/users", func(c *gin.Context) {
			var users []models.User
			// db.Find(&users)
			services.DB.Scopes(services.Paginate(c.Request)).Find(&users)

			log.Println(users)

			c.JSON(http.StatusOK, gin.H{
				"message": "User",
				"payload": users,
			})
		})
			r.POST("/users", func(c *gin.Context) {
				var webUser models.User
				c.Bind(&webUser)
				services.DB.Create(&webUser)
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
					"webUser": webUser,
				})
			})

			r.GET("/users/:id", func(c *gin.Context) {
				var user models.User
				services.DB.First(&user, c.Param("id"))
				services.DB.Find(&user)
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
					"payload": user,
				})
			})

	*/

	r.Run("localhost:8080") //Making sure localhost is there prevents the annoying mac firewall pop up
}
