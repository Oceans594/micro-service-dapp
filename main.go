package main

import (
	"github.com/gin-gonic/gin"

	"micro-service/apps/user"
)

func main() {
	router := gin.Default()
	// 用户组路由
	userRouter := router.Group("api/user")
	{
		userRouter.GET("/login", user.LoginEndpoint)
	}

	router.Run()
}
