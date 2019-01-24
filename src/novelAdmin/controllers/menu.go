package controllers

import (
	"gopkg.in/mgo.v2/bson"
	"novelAdmin/models"
)

type MenuController struct {
	PublicController
}

// 菜单栏列表
func (this *MenuController) Get() {

	menuModel := &models.Menu{}

	menu := menuModel.GetAllMenu(0)

	this.Data["Menu"] = menu

	this.TplName = "menu/index.html"
}

// 修改
func (this *MenuController) Post() {
	pid := this.GetString("pid", "")

	if pid == "" {

		id := this.GetString("id", "")

		menu := &models.Menu{}

		if err := this.ParseForm(menu); err != nil {
			this.sendJson(2, "数据更新失败", "", "")
		}

		if id != "" {
			menu.Id = bson.ObjectIdHex(id)
		}

		if menu.Name == "" {
			this.sendJson(2, "请填写菜单名称", "", "")
		}

		if !menu.Save() {
			this.sendJson(2, "数据更新失败", "", "")
		}
	} else {

		menuChild := &models.MenuChild{}

		if err := this.ParseForm(menuChild); err != nil {
			this.sendJson(2, "数据更新失败", "", "")
		}

		index, _ := this.GetInt("index", -1)
		if index == -1 {
			if !menuChild.AddChild(pid) {
				this.sendJson(2, "数据添加失败", "", "")
			}
		} else {
			if !menuChild.UpdateChild(pid, index) {
				this.sendJson(2, "数据修改失败", "", "")
			}
		}
	}

	this.sendJson(1, "操作成功", "", "/menu/index")
}

// 修改页面
func (this *MenuController) EditPanel() {
	id := this.GetString("id", "")
	index, _ := this.GetInt("index", -1)

	if id == "" {
		panic("params is error")
	}

	menuModel := &models.Menu{}

	if id != "" {
		menuModel.GetMenuById(id)
	}

	parentMenu, err := models.GetParentMenu()

	if err != nil {
		panic("get parent menu error")
	}

	child := models.MenuChild{}
	if index != -1 {
		child = menuModel.Child[index]
	}

	this.Data["Menus"] = parentMenu
	this.Data["Info"] = menuModel
	this.Data["Index"] = index
	this.Data["Child"] = child
	this.TplName = "menu/edit.html"

}

//添加页面
func (this *MenuController) AddPanel() {
	parentMenu, err := models.GetParentMenu()

	if err != nil {
		panic("get parent menu error")
	}

	this.Data["Menus"] = parentMenu

	this.TplName = "menu/add.html"
}
