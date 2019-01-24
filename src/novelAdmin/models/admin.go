package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Admin struct {
	Id            bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Username      string        `bson:"username" json:"username"`
	Nickname      string        `bson:"nickname" json:"nickname"`
	Password      string        `bson:"password" json:"-"`
	Roleid        int           `bson:"role_id" json:"role_id"`
	LastLoginTime int           `bson:"last_login_time" json:"last_login_time"`
	LastLoginIp   int           `bson:"last_login_ip" json:"last_login_ip"`
	LastLocation  string        `bson:"last_location" json:"last_location"`
	Menus         string        `bson:"menus" json:"menus"`
	Addtime       int           `bson:"add_time" json:"add_time"`
	Remark        string        `bson:"remark" json:"remark"`
	Status        byte          `bson:"status" json:"status"`
}

func (this *Admin) tableName() (tableName string) {
	tableName = "n_admin"
	return
}

func (this *Admin) CheckAdmin(username, passwd string) (ret bool) {

	collection, s := getDBSession(this.tableName())

	defer s.Close()

	err := collection.Find(bson.M{"username": username}).One(this) // 返回 QuerySeter

	ret = false
	if err != nil {
		return
	}

	return checkPassWord(this.Password, passwd)
}
