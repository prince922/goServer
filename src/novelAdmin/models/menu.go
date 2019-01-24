package models

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Menu struct {
	Id     bson.ObjectId `bson:"_id,omitempty" json:"id" form:"-"`
	Name   string        `bson:"name" json:"name" form:"name"`
	Url    string        `bson:"url" json:"url" form:"url"`
	Sort   int           `bson:"sort" json:"sort" form:"sort"`
	Status int           `bson:"status" json:"status" form:"status"`
	Child  []MenuChild   `bson:"child" json:"child" form:"-"`
}

type MenuChild struct {
	Name   string `bson:"name" json:"name" form:"name"`
	Url    string `bson:"url" json:"url" form:"url"`
	Sort   int    `bson:"sort" json:"sort" form:"sort"`
	Status int    `bson:"status" json:"status" form:"status"`
}

// 菜单映射表名
func (this *Menu) tableName() (tableName string) {
	tableName = "n_menu"
	return
}

func (this *MenuChild) tableName() (tableName string) {
	tableName = "n_menu"
	return
}

// 获取所有菜单
func (this *Menu) GetAllMenu(status int) (menu []*Menu) {
	collection, s := getDBSession(this.tableName())
	defer s.Close()

	where := bson.M{}
	if status != 0 {
		where = bson.M{"status": status}
	}

	err := collection.Find(where).Sort("-sort").All(&menu)

	if err != nil {
		panic("get menu error")
		return
	}
	return
}

func GetParentMenu() (menus []*Menu, err error) {
	menu := &Menu{}
	collection, s := getDBSession(menu.tableName())
	defer s.Close()
	err = collection.Find(nil).Select(bson.M{"_id": 1, "name": 1}).Sort("-sort").All(&menus)
	return
}

// 添加/修改顶级菜单
func (this *Menu) Save() bool {
	collection, s := getDBSession(this.tableName())
	defer s.Close()
	if this.Id.Hex() == "" {
		err := collection.Insert(this)
		if err != nil {
			return false
		}

	} else {
		if err := collection.Update(bson.M{"_id": this.Id}, bson.M{"$set": bson.M{"name": this.Name, "status": this.Status, "sort": this.Sort}}); err != nil {
			return false
		}

	}
	return true
}

// 添加子菜单
func (this *MenuChild) AddChild(id string) bool {

	collection, s := getDBSession(this.tableName())
	defer s.Close()

	err := collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$push": bson.M{"child": this}})

	if err != nil {
		return false
	}

	return true
}

// 修改子菜单
func (this *MenuChild) UpdateChild(id string, index int) bool {
	collection, s := getDBSession(this.tableName())
	defer s.Close()

	childIndex := fmt.Sprintf("child.%v", index)

	err := collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{childIndex: this}})

	if err != nil {
		return false
	}

	return true
}

// 通过ID获取一条信息
func (this *Menu) GetMenuById(id string) {
	collection, s := getDBSession(this.tableName())
	defer s.Close()
	err := collection.FindId(bson.ObjectIdHex(id)).One(this)
	if err != nil {
		panic("get info error")
	}
}
