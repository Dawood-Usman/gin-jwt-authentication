package main

import (
	"github.com/dawood-usman/go-ops/config"
	"github.com/dawood-usman/go-ops/handlers"
	"github.com/dawood-usman/go-ops/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
	config.MigrateDB()
}

func main() {
	r := gin.Default()

	r.POST("/signup", handlers.SignUp)
	r.POST("/login", handlers.Login)
	r.GET("/validate", middlewares.ValidateAuth, handlers.Validate)

	r.Run()
}
