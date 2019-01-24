package models

import "gopkg.in/mgo.v2/bson"

/**
分类model
*/

type Type struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id" form:"-"`
	Name      string        `bson:"name" json:"name" form:"name"`
	Alias     string        `bson:"alias" json:"alias" form:"alias"`
	Recommend int           `bson:"recommend" json:"recommend" form:"recommend"`
	Channel   int           `bson:"channel" json:"channel" form:"channel"`
	Label     string        `bson:"label" json:"label" form:"label"`
	Sort      int           `bson:"sort" json:"sort" form:"sort"`
}

type SearchTypeListParams struct {
	Name      string `form:"name"`
	Channel   int    `form:"channel"`
	Recommend int    `form:"recommend"`
	Page      int    `form:"page"`
	Size      int    `form:"size"`
}

// 菜单映射表名
func (this *Type) tableName() (tableName string) {
	tableName = "n_type"
	return
}

// 按条件获取分类信息
func (this *Type) GetTypeList(params *SearchTypeListParams) (lists []*Type, count int, err error) {

	collection, s := getDBSession(this.tableName())
	defer s.Close()

	where := bson.M{}

	if params.Name != "" {
		where["name"] = params.Name
	}

	if params.Channel != 0 {
		where["channel"] = params.Channel
	}

	if params.Recommend != 0 {
		where["recommend"] = params.Recommend

	}

	limitStart := (params.Page - 1) * params.Size

	err = collection.Find(where).Sort("-sort").Skip(limitStart).Limit(params.Size).All(&lists)

	// 获取总条数

	count, err = collection.Find(where).Count()

	return
}

// 通过ID获取分类信息
func (this *Type) GetById(id string) {
	collection, s := getDBSession(this.tableName())
	defer s.Close()
	err := collection.FindId(bson.ObjectIdHex(id)).One(this)
	if err != nil {
		panic("get info error")
	}
}

// 获取所有分类
func (this *Type) GetAll() (list []*Type) {

	collection, s := getDBSession(this.tableName())
	defer s.Close()

	_ = collection.Find(bson.M{}).All(&list)
	return
}

func (this *Type) Save() bool {
	collection, s := getDBSession(this.tableName())
	defer s.Close()
	if this.Id.Hex() == "" {
		err := collection.Insert(this)
		if err != nil {
			return false
		}

	} else {
		if err := collection.Update(bson.M{"_id": this.Id}, bson.M{"$set": this}); err != nil {
			return false
		}

	}
	return true
}
