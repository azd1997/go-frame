package http

import (
	"github.com/azd1997/go-frame/process/controller"
	"github.com/gin-gonic/gin"
	"strconv"
)


// websocket请求ping，返回pong
func Ws(ctx *gin.Context) {
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	for {
		// 读数据
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}

		if string(msg) == "ping" {
			msg = []byte("pong")
		}

		if string(msg) == "server_time" {
			_, _, respData := controller.GetServerTime()
			serverTime := strconv.FormatInt(respData.ServerTime, 10)
			msg = []byte(serverTime)
		}

		// 写数据
		err = ws.WriteMessage(msgType, msg)
		if err != nil {
			break
		}
	}
}
