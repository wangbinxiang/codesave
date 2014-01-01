package controllers

import (
	"codesave/libs"
	"github.com/astaxie/beego"
)

type LogoutController struct {
	libs.BaseController
}

func (this *LogoutController) Prepare() {
	this.BaseController.Prepare()
	if this.IsLogin != true {
		this.Redirect("/l", 302)
	}
}

func (this *LogoutController) Get() {
	cookieHash := beego.AppConfig.String("cookieHash")
	cookieName := beego.AppConfig.String("cookieName")
	this.DelSession("userinfo")
	this.SetSecureCookie(cookieHash, cookieName, "", -86400)
	this.Redirect("/", 302)
}
