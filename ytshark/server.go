package main

import (
	"ytshark/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.RedirectFixedPath = true
	router.POST("/post", controller.Post)
	router.GET("/json", controller.ReturnJson)
	router.GET("/json2", controller.ReturnJson2)
	router.Any("/any", controller.Any)
	router.POST("/employee", controller.DemoHandler)
	router.Run(":8080")
}
