package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	confObj, err := config.NewConfig("ini", "/Users/prince/Desktop/prince/goServer/conf/logServer.conf")

	if err != nil {
		fmt.Println(err)
		fmt.Println("init config err")
		return
	}

	logPath := confObj.String("log::logPath")

	if len(logPath) <= 0 {
		fmt.Println("logPath is empty")
		return
	}

	fmt.Println(logPath)
}
