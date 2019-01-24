package models

import "gopkg.in/mgo.v2/bson"

/**
分类model
*/

type Book struct {
	Id         bson.ObjectId `bson:"_id,omitempty" json:"id" form:"-"`
	Name       string        `bson:"name" json:"name" form:"name"`                    // 书名
	SourceId   string        `bson:"source_id" json:"source_id" form:"source_id"`     // 采集源ID
	Tid        string        `bson:"tid" json:"tid" form:"tid"`                       // 类型ID
	Channel    int           `bson:"channel" json:"channel" form:"channel"`           // 1:男，2:女 3:全部
	Label      string        `bson:"label" json:"label" form:"label"`                 // 标签
	Intro      string        `bson:"intro" json:"intro" form:"intro"`                 // 简介
	Status     int           `bson:"status" json:"status" form:"status"`              // 状态 1:上架 -1:下架
	Words      int           `bson:"words" json:"words" form:"words"`                 // 总字数
	CNum       int           `bson:"cnum" json:"cnum" form:"cnum"`                    // 章节数
	Thumb      string        `bson:"thumb" json:"thumb" form:"thumb"`                 // 封面
	IsEnd      int           `bson:"is_end" json:"is_end" form:"is_end"`              // 是否完结
	Author     string        `bson:"author" json:"author" form:"author"`              // 作者
	AuthorID   string        `bson:"author_id" json:"author_id" form:"author_id"`     // 作者ID
	AddTime    int           `bson:"add_time" json:"addTime" form:"addTime"`          // 创建时间
	UpdateTime int           `bson:"update_time" json:"updateTime" form:"updateTime"` // 修改时间
}

type BooKSearch struct {
	Page   int    `json:"page" form:"page"`
	Size   int    `json:"size" form:"size"`
	Name   string `json:"name" form:"name"`
	Status int    `json:"status" form:"status"`
	IsEnd  int    `json:"isEnd" form:"isEnd"`
}

// 菜单映射表名
func (this *Book) tableName() (tableName string) {
	tableName = "n_book"
	return
}

// 按条件查询列表
func (this *Book) GetList(params *BooKSearch) (lists []*Book, count int, err error) {
	collection, s := getDBSession(this.tableName())
	defer s.Close()
	where := bson.M{}

	if params.Name != "" {
		where["name"] = params.Name
	}

	if params.Status != 0 {
		where["status"] = params.Status
	}

	if params.IsEnd != 0 {
		where["is_end"] = params.IsEnd

	}

	limitStart := (params.Page - 1) * params.Size

	err = collection.Find(where).Skip(limitStart).Limit(params.Size).All(&lists)
	count, err = collection.Find(where).Count()

	return
}

func (this *Book) Add() bool {
	collection, s := getDBSession(this.tableName())
	defer s.Close()

	err := collection.Insert(this)
	if err != nil {
		return false
	}
	return true
}

func (this *Book) Update() bool {
	collection, s := getDBSession(this.tableName())
	defer s.Close()
	if err := collection.Update(bson.M{"_id": this.Id}, bson.M{"$set": this}); err != nil {
		return false
	}
	return true
}

func (this *Book) GetById(id string) {
	collection, s := getDBSession(this.tableName())
	defer s.Close()
	err := collection.FindId(bson.ObjectIdHex(id)).One(this)
	if err != nil {
		panic("get info error")
	}
}
