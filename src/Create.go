package src

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Mireu-Lab/golang-practice-project/src/csvfile"
	"github.com/gin-gonic/gin"
)

func Create(rg *gin.RouterGroup) {
	create := rg.Group("/create")
	go create.POST("/query", QueryType_Create_Program)
	go create.POST("/json", JSONType_Create_Program)
}

type postJsonInfo struct {
	msg string `json:msg`
}

func QueryType_Create_Program(g *gin.Context) {
	msg := g.Query("msg")
	var WriteData = []string{time.RFC3339, msg}

	err := csvfile.Write("./File/output.csv", WriteData)
	if err != nil {
		log.Fatalln(err)
	}

	g.JSON(http.StatusOK, H{
		"time": time.RFC3339,
		"msg":  msg})
}

func JSONType_Create_Program(g *gin.Context) {
	var Jsondata postJsonInfo

	err := json.NewDecoder(g.Request.Body).Decode(&Jsondata)
	if err != nil {
		log.Fatalln(err)
	}

	var WriteData = []string{time.RFC3339, Jsondata.msg}

	err = csvfile.Write("./File/output.csv", WriteData)
	if err != nil {
		log.Fatalln(err)
	}

	g.JSON(http.StatusOK, H{
		"time": time.RFC3339,
		"msg":  Jsondata})
}
