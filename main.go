package main

import (
	"github.com/Mireu-Lab/golang-practice-project/src"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {
	src.Ping(&router.RouterGroup)

	v1 := router.Group("/v1")
	src.Ping(v1)
	src.Lists(v1)
	src.Create(v1)
	// src.Updata(v1)
	// src.Delete(v1)

	// router.GET("/", src.Lists())
	// router.POST("/", src.Create())
	// router.DELETE("/", src.Delete())

	router.Run(":80")
}
