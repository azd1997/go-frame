package http

import (
	"github.com/azd1997/go-frame/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

var engine *gin.Engine

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
}
