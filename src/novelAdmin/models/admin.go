package models

import ()

type Admin struct {
	Id            int    `orm:"auto" json:"id"`
	Username      string `json:"username"`
	Nickname      string `json:"nickname"`
	Password      string `json:"password"`
	Roleid        int    `json:"roleid"`
	LastLoginTime int    `json:"lastLoginTime"`
	LastLoginIp   int    `json:"lastLoginIp"`
	LastLocation  string `json:"lastLocation"`
	Menus         string `json:"menus"`
	Addtime       int    `json:"addtime"`
	Remark        string `json:"remark"`
	Status        byte   `json:"status"`
}

func (this *Admin) CheckAdmin(username, passwd string) (ret bool) {

	this.Username = username
	db.Using("default")

	query := db.QueryTable(this) // 返回 QuerySeter

	err := query.Filter("username", username).One(this, "Id", "Username", "Password", "Roleid", "Nickname")

	ret = false
	if err != nil {
		return
	}

	return checkPassWord(this.Password, passwd)
}
