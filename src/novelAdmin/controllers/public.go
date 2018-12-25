package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

type RetJson struct {
	code int
	msg  string
	data interface{}
}

type PublicController struct {
	beego.Controller
}

func (this *PublicController) sendJson(code int, msg string, data interface{}) {
	ret := &RetJson{1, "", ""}
	ret.code = code
	ret.msg = msg
	ret.data = data
	this.Data["json"] = ret
	this.ServeJSON()
	this.StopRun()
}

// 初始化redis
func getRedis() (redis cache.Cache) {
	host := beego.AppConfig.String("RedisHost")
	port := beego.AppConfig.String("RedisPort")

	redis, err := cache.NewCache("redis", `{"key":"novel","conn":"`+host+`:`+port+`","dbNum":"0"}`)

	if err != nil {
		beego.Info("init redis error")
		panic("init redis error")
		return
	}
	return
}

// 调试
func showDebug(msg interface{}) {
	beego.Debug(msg)
}
