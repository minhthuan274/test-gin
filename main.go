package main

import (
	"github.com/gin-gonic/gin"
	"github.com/minhthuan274/test-gin/db"
	"github.com/minhthuan274/test-gin/handlers/users"
)

func init() {
	db.Connect()
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	v3 := r.Group("/api/v3")
	{
		v3.GET("/users", users.fetchAllUsers)
		v3.GET("/users/:id", users.fetchUser)
	}

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
