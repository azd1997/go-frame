// 程序入口文件
package main

import (
	"flag"
	"fmt"
	"github.com/azd1997/go-frame/cache/redis"
	"github.com/azd1997/go-frame/config"
	"github.com/azd1997/go-frame/db"
	"github.com/azd1997/go-frame/logger"
	"github.com/azd1997/go-frame/mq/nsq"
	"github.com/azd1997/go-frame/process/http"
	"github.com/azd1997/go-frame/process/rpc"
	"os"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "配置文件路径")
	flag.Parse()

	if configPath == "" {
		fmt.Println("Config path must be assigned.")
		os.Exit(1)
	}

	fmt.Println(configPath)

	var err error


	// 加载配置
	err = config.InitConfig(configPath)
	if err != nil {
		fmt.Printf("InitConfig failed: %v\n", err)
		os.Exit(1)
	}

	// 开启日志记录
	err = logger.InitLogger(config.Conf().LoggerConfig)
	if err != nil {
		fmt.Printf("InitLogger failed: %v\n", err)
		os.Exit(1)
	}

	// 开启mysql数据库连接
	err = db.InitEngine(config.Conf().DBConfig)
	if err != nil {
		fmt.Printf("InitEngine failed: %v\n", err)
		os.Exit(1)
	}

	// 开启redis缓存连接
	err = redis.InitRedis(config.Conf().RedisConfig)
	if err != nil {
		fmt.Printf("InitRedis failed: %v\n", err)
		os.Exit(1)
	}

	// 启动HTTP服务
	go http.StartHTTPServer(config.Conf().HttpConfig)

	// 启动RPC服务
	go rpc.StartRPCServer(config.Conf().RpcConfig)

	// 启动消息队列
	go nsq.StartNsqServer(config.Conf().NsqConfig)

	logger.Logger().Info("Init success!")

	// 阻止主go程退出
	select {}
}