package main

import (
	"go-restapi-jwt-gin-postgresql/controllers"
	"go-restapi-jwt-gin-postgresql/middlewares"
	"go-restapi-jwt-gin-postgresql/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	// public.POST("/register", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})
	// })

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
