package main

import (
	"log"
	"wmw-user-api/middleware"
	"wmw-user-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/login", services.User.Login)
	r.POST("/register", services.User.Register)
	rAuth := r.Group("/")
	rAuth.Use(middleware.Auth())
	rAuth.GET("/info", services.User.Info)
	log.Fatal(r.Run("127.0.0.1:8080"))
}
