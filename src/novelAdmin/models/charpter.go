package models

import "gopkg.in/mgo.v2/bson"

/**
分类model
*/

type Charpter struct {
	Id         bson.ObjectId `bson:"_id,omitempty" json:"id" form:"-"`
	Bid        string        `bson:"bid" json:"bid" form:"bid"`                       // 小说ID
	SourceId   string        `bson:"source_id" json:"source_id" form:"source_id"`     // 采集源小说ID
	SourceCid  string        `bson:"source_cid" json:"source_cid" form:"source_cid"`  // 采集源小说篇节ID
	Name       string        `bson:"name" json:"name" form:"name"`                    // 小说卷名
	Content    string        `bson:"content" json:"content" form:"content"`           // 章节内容
	Status     int           `bson:"status" json:"status" form:"status"`              // 状态 1:上架 -1:下架
	Sort       int           `bson:"sort" json:"sort" form:"sort"`                    // 排序
	AddTime    int           `bson:"add_time" json:"addTime" form:"addTime"`          // 创建时间
	UpdateTime int           `bson:"update_time" json:"updateTime" form:"updateTime"` // 修改时间
}

// 菜单映射表名
func (this *Charpter) tableName() (tableName string) {
	tableName = "n_charpter"
	return
}

// 按条件查询列表
func (this *Charpter) GetList(params *BooKSearch) (lists []*Book, count int, err error) {
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

func (this *Charpter) Add() bool {
	collection, s := getDBSession(this.tableName())
	defer s.Close()

	err := collection.Insert(this)
	if err != nil {
		return false
	}
	return true
}

func (this *Charpter) Update() bool {
	collection, s := getDBSession(this.tableName())
	defer s.Close()
	if err := collection.Update(bson.M{"_id": this.Id}, bson.M{"$set": this}); err != nil {
		return false
	}
	return true
}

func (this *Charpter) GetById(id string) {
	collection, s := getDBSession(this.tableName())
	defer s.Close()
	err := collection.FindId(bson.ObjectIdHex(id)).One(this)
	if err != nil {
		panic("get info error")
	}
}
