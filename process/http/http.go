package http

import (
	"github.com/azd1997/go-frame/config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"os"
)

var engine *gin.Engine

// 用于将http升级为websocket
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 启动http服务
func StartHTTPServer(conf config.HttpConfig) {
	engine = gin.Default()
	Route()
	err := engine.Run(conf.Addr)
	if err != nil {
		zap.Error(err)
		os.Exit(1)
	}
}

// 路由
func Route() {
	engine.GET("/server_time", GetServerTime)
	engine.GET("/ws", Ws)
}
