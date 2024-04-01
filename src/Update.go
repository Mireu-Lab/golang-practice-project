package src

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Updata(rg *gin.RouterGroup) {
	updata := rg.Group("updata")
	updata.PUT("/", Updata_program)
}

func Updata_program(g *gin.Context) {
	g.JSON(200, H{"time": time.RFC3339})
}
