package controllers

import (
	"gopkg.in/mgo.v2/bson"
	"novelAdmin/comment"
	"novelAdmin/models"
)

type TypeController struct {
	PublicController
}

// 栏列表
func (this *TypeController) Get() {

	FormParams := &models.SearchTypeListParams{}
	if err := this.ParseForm(FormParams); err != nil {
		panic("参数出错")
	}

	if FormParams.Size == 0 {
		FormParams.Size = 10
	}
	if FormParams.Page == 0 {
		FormParams.Page = 1
	}

	model := &models.Type{}

	lists, count, err := model.GetTypeList(FormParams)

	if err != nil {
		panic("参取数据失败")
		return
	}

	pageModel := &comment.Page{count, int(FormParams.Size)}

	pageInfo := pageModel.ShowPage(int(FormParams.Page), 10)

	this.Data["Lists"] = lists
	this.Data["Page"] = Struct2Json(pageInfo)
	this.Data["Form"] = FormParams
	this.TplName = "type/index.html"
}

// 修改
func (this *TypeController) Post() {
	id := this.GetString("id", "")

	if id == "" {
		panic("params is error")
	}

	model := &models.Type{}

	if err := this.ParseForm(model); err != nil {
		this.sendJson(2, "数据更新失败", "", "")
	}

	model.Id = bson.ObjectIdHex(id)
	if !model.Save() {
		this.sendJson(2, "数据更新失败", "", "")
	}

	this.sendJson(1, "操作成功", "", "/type/index")
}

// 添加页面
func (this *TypeController) EditPanel() {
	id := this.GetString("id", "")

	model := &models.Type{}

	if id != "" {
		model.GetById(id)
	}

	parentMenu, err := models.GetParentMenu()

	if err != nil {
		panic("get parent menu error")
	}

	this.Data["Menus"] = parentMenu
	this.Data["Info"] = model
	this.TplName = "type/edit.html"

}
