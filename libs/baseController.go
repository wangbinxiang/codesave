package libs

import (
	"github.com/astaxie/beego"
	m "github.com/wangbinxiang/codesave/models"
	"strings"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	LoginUser m.UserAccount
}

func (this *BaseController) Prepare() {
	this.checkLogin()
	this.InitHtml()
}

func (this *BaseController) Finish() {
}

func (this *BaseController) InitHtml() {
	this.Layout = "layout.html"
	this.LayoutSections = make(map[string]string)
}

func (this *BaseController) checkLogin() {
	this.IsLogin = false
	cookieHash := beego.AppConfig.String("cookieHash")
	cookieName := beego.AppConfig.String("cookieName")
	cookieSep := beego.AppConfig.String("cookieSep")
	userAccount := this.GetSession("userinfo")
	if userAccount != nil {
		this.IsLogin = true
		this.LoginUser = userAccount.(m.UserAccount)
	} else {
		userCookie, _ := this.GetSecureCookie(cookieHash, cookieName)
		parts := strings.Split(userCookie, cookieSep)
		if len(parts) == 3 {
			email := parts[1]
			password := parts[2]
			userAccount, err := m.GetUserAccountByEmail(email)

			if err == nil {
				//检查密码
				if userAccount.Password == password {
					this.IsLogin = true
					this.LoginUser = userAccount
					this.SetSession("userinfo", userAccount)
					//设置session Cookie
				} else {
					this.SetSecureCookie(cookieHash, cookieName, "", -86400)
				}
			}
		}
	}

	if this.IsLogin {
		this.Data["isLogin"] = true
		this.Data["user"] = this.LoginUser
	} else {
		this.Data["isLogin"] = false
	}
}

func (this *BaseController) LoginJump(login bool) { //登陆跳转 true:要求登陆 false:要求不登陆
	if login {
		if !this.IsLogin {
			if this.IsAjax() {
				this.StopRun()
			} else {
				this.Redirect("/", 302)
			}
		}
	} else {
		if this.IsLogin {
			this.Redirect("/", 302)
		}
	}

}
