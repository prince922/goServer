package controllers

import (
	"gopkg.in/mgo.v2/bson"
	"novelAdmin/comment"
	"novelAdmin/models"
	"strings"
)

type CharpterController struct {
	PublicController
}

// 栏列表
func (this *CharpterController) Get() {

	search := &models.BooKSearch{}
	if err := this.ParseForm(search); err != nil {
		panic("参数出错")
	}

	if search.Size == 0 {
		search.Size = 10
	}
	if search.Page == 0 {
		search.Page = 1
	}

	model := &models.Book{}

	lists, count, err := model.GetList(search)

	if err != nil {
		panic("参取数据失败")
		return
	}

	pageModel := &comment.Page{count, int(search.Size)}

	pageInfo := pageModel.ShowPage(int(search.Page), 10)

	this.Data["Lists"] = lists
	this.Data["Page"] = Struct2Json(pageInfo)
	this.Data["Form"] = search
	this.TplName = "book/index.html"
}

func (this *CharpterController) Post() {
	id := this.GetString("id", "")

	model := &models.Book{}

	if err := this.ParseForm(model); err != nil {
		this.Redirect("/book/index", 302)
	}

	if id == "" {

		model.Id = bson.NewObjectId()

		savePath := "book"
		fileName := model.Id.Hex()

		filePath, _ := this.saveImage("thumb", savePath, fileName)

		model.Thumb = filePath

		if !model.Add() {
			this.Redirect("/book/index", 302)
			//this.sendJson(2,"数据添加失败", "", "")
		}

	} else {
		model.Id = bson.ObjectIdHex(id)

		if !model.Update() {
			this.Redirect("/book/index", 302)
		}

	}
	this.Redirect("/book/index", 302)

}

// 添加页面
func (this *CharpterController) EditPanel() {
	id := this.GetString("id", "")

	model := &models.Book{}

	if id != "" {
		model.GetById(id)
	}

	typeModel := &models.Type{}

	typeList := typeModel.GetAll()

	labelArr := strings.Split(model.Label, ",")

	this.Data["Info"] = model
	this.Data["LabelArr"] = labelArr
	this.Data["TypeList"] = typeList
	this.TplName = "book/edit.html"

}
