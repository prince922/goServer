package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"novelAdmin/comment"
)

var (
	url    string
	dbName string
)

func init() {
	// 从配置文件读信息
	url = beego.AppConfig.String("MonUrl")
	dbName = beego.AppConfig.String("MonName")

}

func getDBSession(tableName string) (c *mgo.Collection, session *mgo.Session) {
	session, err := mgo.Dial(url)

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	if err != nil {
		panic(err)
	}
	c = session.DB(dbName).C(tableName)
	return
}

// 对比密码
func checkPassWord(dbPw, postPw string) bool {
	return createPassWord(postPw) == dbPw
}

// 生成密码
func createPassWord(passWord string) string {
	return comment.Str2md5(passWord + "novel")
}

// 调试
func showDebug(msg interface{}) {
	beego.Debug(msg)
}
