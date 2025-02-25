package main

import (
	"github.com/gin-gonic/gin"
	"task/database"
	"task/middlewares"
	"task/routes"
)

func main() {

	database.InitDB()

	g := gin.Default()

	g.GET("/tasks", middlewares.RequireAuth, database.GetHandler)
	g.GET("/validate", middlewares.RequireAuth, routes.Validate)

	g.POST("/tasks", middlewares.RequireAuth, database.PostHandler)
	g.POST("/signup", routes.Signup)
	g.POST("/login", routes.Login)

	g.DELETE("/tasks/:id", middlewares.RequireAuth, database.DeleteHandler)

	g.PATCH("/tasks/:id", middlewares.RequireAuth, database.PatchHandler)

	g.Run(":8080")
}
