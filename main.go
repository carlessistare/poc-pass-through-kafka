package main

import "fmt"
import "github.com/gin-gonic/gin"

import "ws-test/controllers"

func main() {
	fmt.Println("Hello, 世界")
	r := gin.Default()

	r.GET("/ping", controllers.Ping)
	r.POST("/proto", controllers.Proto)

	r.Run(":8001")
}
