package src

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Mireu-Lab/golang-practice-project/src/sql"
	"github.com/gin-gonic/gin"
)

func Create(rg *gin.RouterGroup) {
	create := rg.Group("/create")
	go create.POST("/query", QueryType_Create_Program)
	go create.POST("/json", JSONType_Create_Program)
}

func QueryType_Create_Program(g *gin.Context) {
	msg := g.Query("msg")

	errCode := sql.Write(msg)

	g.JSON(errCode, H{
		"time": time.RFC3339,
		"msg":  msg})
}

type postJsonInfo struct {
	msg string `json:msg`
}

func JSONType_Create_Program(g *gin.Context) {
	var errCode int = 500
	var input postJsonInfo
	var ReturnJson = H{
		"path": "/v1/json",
		"time": time.RFC3339,
		"msg":  input.msg}

	err := json.NewDecoder(g.Request.Body).Decode(&input)
	if err != nil {
		log.Fatalln(err)
	}

	// err := g.ShouldBindBodyWith(&Jsondata, binding.JSON)
	// err := g.ShouldBindJSON(&input)
	// if err != nil {
	// 	errCode = 403
	// 	ReturnJson["ErrorMessage"] = err.Error()
	// }

	fmt.Println(input)

	if input.msg != "" {
		errCode = sql.Write(input.msg)
	} else {
		ReturnJson["ErrorMessage"] = "The message is empty."
	}

	g.JSON(errCode, ReturnJson)
}
