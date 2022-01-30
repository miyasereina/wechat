package router

import (
	"github.com/gin-gonic/gin"
	"wechat/controller"
)

func SetRouter() *gin.Engine{
	r:=gin.Default()
	r.GET("/echostr",controller.CheckEchostr)
	return r
}
