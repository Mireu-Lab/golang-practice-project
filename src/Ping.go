package src

import (
	"time"

	"github.com/gin-gonic/gin"
)

// gin setting
func Ping(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")
	go ping.GET("/", Ping_Program)
}

func Ping_Program(g *gin.Context) {
	g.JSONP(200, H{
		"time": time.RFC3339,
		"msg":  "pong"})
}
