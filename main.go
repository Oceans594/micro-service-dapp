package main

import (
	"github.com/gin-gonic/gin"

	"micro-service/apps/application"
	"micro-service/apps/user"
)

func main() {
	router := gin.Default()
	// 用户组路由
	userRouter := router.Group("api/user")
	{
		userRouter.GET("/login", user.LoginEndpoint)
	}
	// 应用组路由
	applicationRouter := router.Group("api/app")
	{
		applicationRouter.GET("/detail", application.DetailEndpoint)
	}

	router.Run()
}
