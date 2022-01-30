package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

const token = "riribaka"

type CheckMsg struct {
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
	Nonce string `json:"nonce"`
	Echostr string `json:"echostr"`
}

func CheckEchostr(ctx *gin.Context)  {
	checkMsg:=CheckMsg{}
	err:=ctx.ShouldBindJSON(&checkMsg)
	if err != nil {
		log.Println(err)
	}
	//tmpArr:= map[string]{token,checkMsg.Timestamp,checkMsg.Nonce}
	ctx.JSON(200,gin.H{
		"echostr":checkMsg.Echostr,
	})
}
