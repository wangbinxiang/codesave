package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wangbinxiang/codesave/libs"
)

type LogoutController struct {
	libs.BaseController
}

func (this *LogoutController) Prepare() {
	this.BaseController.Prepare()
	this.LoginJump(true)
}

func (this *LogoutController) Get() {
	cookieHash := beego.AppConfig.String("cookieHash")
	cookieName := beego.AppConfig.String("cookieName")
	this.DelSession("userinfo")
	this.SetSecureCookie(cookieHash, cookieName, "", -86400)
	this.Redirect("/", 302)
}
