package http

import (
	"github.com/azd1997/go-frame/process/controller"
	"github.com/gin-gonic/gin"
)

type Request struct {}

type Response struct {
	Code int	`json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func GetServerTime(ctx *gin.Context) {
	resp := Response{}
	resp.Code, resp.Msg, resp.Data = controller.GetServerTime()
	ctx.JSON(resp.Code, resp)
}


