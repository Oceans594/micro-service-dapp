package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func LoginEndpoint(c *gin.Context) {
	// 调用Python 用户微服务
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c.String(http.StatusOK, "Login Success")
}
