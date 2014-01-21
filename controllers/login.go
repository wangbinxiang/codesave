package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	h "github.com/wangbinxiang/codesave/helper"
	"github.com/wangbinxiang/codesave/libs"
	m "github.com/wangbinxiang/codesave/models"
	"html/template"
	"log"
	"strconv"
)

type LoginController struct {
	libs.BaseController
}

func (this *LoginController) Prepare() {
	this.BaseController.Prepare()
	this.LoginJump(false)
}

func (this *LoginController) Get() {

	this.LayoutSections["htmlFooter"] = "footer/loginFooter.html"

	this.Data["xsrfdata"] = template.HTML(this.XsrfFormHtml())

	this.TplNames = "templates/login.html"
}

func (this *LoginController) Post() {
	email := this.GetString("Email")
	password := this.GetString("Password")

	valid := validation.Validation{}
	valid.Email(email, "email")
	valid.MaxSize(email, 128, "email")
	valid.MinSize(password, 6, "password")
	valid.MaxSize(password, 16, "password")

	// [\x00-\xff] 密码输入正则
	//检查 输入
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err)
			this.Redirect("/l", 302)
		}
	}

	//获取对应用户
	userAccount, err := m.GetUserAccountByEmail(email)

	if err != nil {
		this.Redirect("/l", 302)
	} else {
		//检查密码
		if userAccount.Password == h.MD5(password+userAccount.Salt) {
			cookieHash := beego.AppConfig.String("cookieHash")
			cookieName := beego.AppConfig.String("cookieName")
			cookieSep := beego.AppConfig.String("cookieSep")

			cookieStr := strconv.Itoa(int(userAccount.Id)) + cookieSep + userAccount.Email + cookieSep + userAccount.Password
			this.SetSecureCookie(cookieHash, cookieName, cookieStr, 604800)
			this.SetSession("userinfo", userAccount)
			//设置session Cookie
			this.Redirect("/", 302)
		} else {
			this.Redirect("/l", 302)
		}
	}
}
