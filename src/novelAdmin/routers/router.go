package routers

import (
	"github.com/astaxie/beego"
	"novelAdmin/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/login", &controllers.LoginController{})
	beego.Router("/login/logout", &controllers.LoginController{}, "get:Logout")
}
