package routers

import (
	"github.com/astaxie/beego"
	"novelAdmin/controllers"
)

func init() {

	// 后台首页
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{}, "get:Index")

	// 后台登录页
	beego.Router("/admin/login", &controllers.LoginController{})
	beego.Router("/login/logout", &controllers.LoginController{}, "get:Logout")

	// 菜单栏
	beego.Router("/menu/index", &controllers.MenuController{})
	beego.Router("/menu/edit", &controllers.MenuController{}, "get:EditPanel;post:Post")
	beego.Router("/menu/add", &controllers.MenuController{}, "get:AddPanel")

	// 小说管理
	beego.Router("/book/index", &controllers.BookController{})
	beego.Router("/book/edit", &controllers.BookController{}, "get:EditPanel;post:Post")

	// 小说章节
	beego.Router("/charpter/index", &controllers.CharpterController{})
	beego.Router("/charpter/edit", &controllers.CharpterController{}, "get:EditPanel;post:Post")

	// 分类管理
	beego.Router("/type/index", &controllers.TypeController{})
	beego.Router("/type/edit", &controllers.TypeController{}, "get:EditPanel;post:Post")
}
