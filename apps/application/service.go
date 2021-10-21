package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"micro-service/apps/application/service/micro-service/apps/application/pb"
)

func DetailEndpoint(c *gin.Context) {
	// 调用python数据服务
	identifier := c.DefaultQuery("identifier", "")
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataCenterClient(conn)
	resp, err := client.GetApplication(context.Background(), &pb.GetApplicationRequest{Identifier: identifier})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Message)
	c.String(http.StatusOK, resp.Message)
}
