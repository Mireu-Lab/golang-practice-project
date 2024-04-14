package src

import (
	"encoding/json"
	"log"
	"time"

	"github.com/Mireu-Lab/golang-practice-project/src/csvfile"
	"github.com/gin-gonic/gin"
)

func Create(rg *gin.RouterGroup) {
	create := rg.Group("/create")
	go create.POST("/query", QueryType_Create_Program)
	go create.POST("/json", JSONType_Create_Program)
}

type CreateJsonInfo struct {
	msg string `json:msg`
}

func QueryType_Create_Program(g *gin.Context) {
	msg := g.Query("msg")
	var WriteData = []string{time.RFC3339, msg}

	err := csvfile.Write("./File/output.csv", WriteData)
	if err != nil {
		log.Fatalln(err)
	}

	g.JSON(200, H{
		"time": time.RFC3339,
		"msg":  msg})
}

func JSONType_Create_Program(g *gin.Context) {
	var CreateJsondata CreateJsonInfo

	err := json.NewDecoder(g.Request.Body).Decode(&CreateJsondata)
	if err != nil {
		log.Fatalln(err)
	}

	var WriteData = []string{time.RFC3339, CreateJsondata.msg}

	err = csvfile.Write("./File/output.csv", WriteData)
	if err != nil {
		log.Fatalln(err)
	}

	g.JSON(200, H{
		"time": time.RFC3339,
		"msg":  CreateJsondata})
}
