package controllers

import (
	"encoding/json"
	"fmt"
	"novelAdmin/models"
	"time"
	//"github.com/astaxie/beego"
)

type LoginController struct {
	PublicController
}

func (this *LoginController) Get() {
	this.TplName = "login/index.html"
}

// 管理登录
func (this *LoginController) Post() {

	userName := this.GetString("username")
	passwd := this.GetString("password")

	admin := &models.Admin{}

	isLogin := admin.CheckAdmin(userName, passwd)

	// 管理登录成功
	if isLogin == true {
		// 把相关用户信息存储起来
		var saveUserInfo = func(info *models.Admin) {
			redis := getRedis()
			key := fmt.Sprintf("adminInfo_%v", info.Id.Hex())
			value, err := json.Marshal(info)
			if err != nil {
				return
			}
			redis.Put(key, string(value), 86400*time.Second)
		}

		// 异步协程处理用户信息
		go saveUserInfo(admin)
		this.SetSession("uid", admin.Id.Hex())
		this.Redirect("/", 302)
		return
	}

	this.Data["msg"] = "登录失败"
	this.TplName = "login/index.html"
}

// 退出登录
func (this *LoginController) Logout() {
	this.DelSession("uid")
	this.Redirect("/admin/login", 302)
}
