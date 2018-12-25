package main

import (
	"github.com/astaxie/beego/context"
	//"github.com/astaxie/beego"
)

// 过滤器
var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)

	// 没有登录，返回到登录
	if !ok && ctx.Request.RequestURI != "/admin/login" {
		ctx.Redirect(302, "/admin/login")
	}

	// 登录后，打开登录页，跳到首页
	if ok && ctx.Request.RequestURI == "/admin/login" {
		ctx.Redirect(302, "/")
	}
}

// 修改 html 方法
var FilterMethod = func(ctx *context.Context) {
	if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
		ctx.Request.Method = ctx.Input.Query("_method")
	}
}
