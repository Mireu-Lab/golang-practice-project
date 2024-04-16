package src

import (
	"time"

	"github.com/Mireu-Lab/golang-practice-project/src/sql"
	"github.com/gin-gonic/gin"
)

func Lists(rg *gin.RouterGroup) {
	ping := rg.Group("/lists")
	go ping.GET("/", Lists_Program)
}

func Lists_Program(g *gin.Context) {
	g.JSONP(200, H{
		"cell_time": time.RFC3339,
		"msg_list":  sql.Read()})
}
