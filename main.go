package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/getuser")
	r.GET("/getuser")

	r.Run(":2323")
}