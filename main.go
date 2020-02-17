// 程序入口文件
package main

import (
	"flag"
	"fmt"
	"github.com/azd1997/go-frame/config"
	"github.com/azd1997/go-frame/logger"
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

	loggerConfig := config.GetConfig().LoggerConfig

	// 开启日志记录
	err = logger.InitLogger(loggerConfig.LogPath, loggerConfig.LogLevel)
	if err != nil {
		fmt.Printf("InitLogger failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(loggerConfig)


	logger.GetLogger().Info("Init success!")
}