package controllers

import (
	"encoding/json"
	"fmt"
	"novelAdmin/models"
)

type MainController struct {
	PublicController
}

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
	this.Data["nickname"] = admin.Nickname
	this.TplName = "index.html"
}
