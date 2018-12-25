package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"novelAdmin/comment"
)

var (
	db orm.Ormer
)

func init() {

	dbhost := beego.AppConfig.String("DBHost")
	dbport := beego.AppConfig.String("DBPort")
	dbuser := beego.AppConfig.String("DBUser")
	dbpasswd := beego.AppConfig.String("DBPw")
	dbName := beego.AppConfig.String("DBName")

	// 注册mysql Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 构造conn连接
	// 用户名:密码@tcp(url地址)/数据库
	conn := dbuser + ":" + dbpasswd + "@tcp(" + dbhost + ":" + dbport + ")/" + dbName + "?charset=utf8"

	// 注册数据库连接
	orm.RegisterDataBase("default", "mysql", conn)
	// 设置数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 30)
	// 设置数据库的最大数据库连接
	orm.SetMaxOpenConns("default", 60)

	// 需要在init中注册定义的model
	regisertModel()
	orm.Debug = true

	db = orm.NewOrm()
}

func regisertModel() {
	orm.RegisterModelWithPrefix("go_", new(Admin))
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
