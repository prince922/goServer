package config

import (
	"github.com/astaxie/beego/config"
)

// 全局变量
var (
	appConfig *configServer
)

type configServer struct {
	LogPath  string
	LogLevel string

	Collect []collect
}

type collect struct {
	Topic string
	Path  string
}

func initCollect(confObj config.Configer) (cl collect, err error) {
	var cl = *collect{}

	cl.Topic = confObj.String("collect::topic")

	if len(cl.Topic) <= 0 {
		panic("read collect topic error")
		return
	}

	cl.Path = confObj.String("collect::path")

	if len(cl.Path) <= 0 {
		panic("read collect path error")
		return
	}

	return
}

// 配置文件初始化
func initConfig(configType, configPath string) (err error) {
	configObj, err := config.NewConfig(configType, configPath)
	if err != nil {
		panic("init config err")
		return
	}

	appConfig = *configServer{}
	// 读取日志路径及级别
	appConfig.LogPath = configObj.string("log::logPath")
	appConfig.LogLevel = configObj.string("log::logLevel")

	cl, err := initCollect(configObj)

	return
}
