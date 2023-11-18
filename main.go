package main

import (
	"context"
	"wmw-user-api/middleware"
	"wmw-user-api/services"

	"github.com/golang/glog"

	"github.com/fxh111111/utility/config"
	"github.com/fxh111111/utility/config/redis"
	"github.com/fxh111111/utility/mongo"
	"github.com/gin-gonic/gin"
	goredis "github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.TODO()
	config.SetKVProvider(redis.GetRedisKV(goredis.NewClient(&goredis.Options{Addr: ":6379"})))
	if err := mongo.Connect(config.GetString(ctx, "mongo:connection", "")); err != nil {
		glog.Fatal(err)
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/login", services.User.Login)
	r.POST("/register", services.User.Register)
	rAuth := r.Group("/")
	rAuth.Use(middleware.Auth())
	rAuth.GET("/info", services.User.Info)
	glog.Fatal(r.Run("127.0.0.1:8080"))
}
