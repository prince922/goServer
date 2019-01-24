package controllers

import (
	"encoding/json"
	"fmt"
	"novelAdmin/models"
	"runtime"
)

type MainController struct {
	PublicController
}

type mainInfo struct {
	Name string
	Info string
}

// 登录后首页
func (this *MainController) Get() {
	adminId := this.GetSession("uid")
	redis := getRedis()

	key := fmt.Sprintf("adminInfo_%v", adminId)

	value := redis.Get(key).([]byte)
	var admin models.Admin
	err := json.Unmarshal(value, &admin)
	if err != nil {
		panic("init admin error")
	}

	// 菜单

	menuModel := &models.Menu{}
	menu := menuModel.GetAllMenu(1)

	this.Data["Menu"] = menu
	this.Data["nickname"] = admin.Nickname
	this.TplName = "default/index.html"
}

// 登录后首页
func (this *MainController) Index() {
	var info []*mainInfo

	info = append(info, &mainInfo{"操作系统", runtime.GOOS})
	info = append(info, &mainInfo{"GO版本", runtime.Version()})

	this.Data["info"] = info
	this.TplName = "default/main.html"
}
