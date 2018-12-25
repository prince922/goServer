package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis" // session redis
	_ "novelAdmin/routers"                     // 路由
)

// 初始化
func init() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "novelappid"
	beego.BConfig.WebConfig.Session.SessionProvider = beego.AppConfig.String("SessionProvider")
	beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("SessionProviderConfig")
}

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/*", beego.BeforeRouter, FilterMethod)

	beego.Run()

}
