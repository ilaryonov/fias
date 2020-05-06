package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_"fmt"
	_"time"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc"
	pb "learning/grpcservice/grpc"
	"golang.org/x/net/context"
	_"fmt"
	"log"
)

type json struct {
	text []string
}

var Client pb.GrpcClient

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		var message []string
		for i:=0;i<1000;i++ {
			response := sendGrpcRequest(Client)
			message = append(message, response)
		}
		log.Println("catch")
		c.JSON(http.StatusOK, gin.H{"text": message})
	})

	return r
}

func main() {
	/*db, err := gorm.Open("mysql", "user:pass@tcp(172.17.0.3)/go_gin?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("db err := ",err)
	}
	db.AutoMigrate(&Product{})

	// Add index for columns `title` with given name `idx_title`
	db.Model(&Product{}).AddIndex("idx_title", "title")*/

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:5300", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	Client = pb.NewGrpcClient(conn)

	r := setupRouter()
	r.Run(":8083")
}

func sendGrpcRequest(client pb.GrpcClient) string {

	request := &pb.Request{
		Message: "Hello World",
	}
	response, err := client.GetString(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	return response.Message
}
