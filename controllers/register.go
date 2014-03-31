package controllers

import (
	"github.com/astaxie/beego"
	. "github.com/wangbinxiang/codesave/caches/memcache"
	h "github.com/wangbinxiang/codesave/helper"
	"github.com/wangbinxiang/codesave/libs"
	m "github.com/wangbinxiang/codesave/models"
	"html/template"
	"log"
	"strconv"
	// "time"
)

const (
	RegisterKeyPrefix       string = "rKey_"
	RegisterExpired         int64  = 3600
	RegisterRecaptchaNumber        = 3
)

type RegisterController struct {
	libs.BaseController
}

func (this *RegisterController) Prepare() {
	this.BaseController.Prepare()
	this.LoginJump(false)
}

func (this *RegisterController) Get() {
	registerKey := RegisterKeyPrefix + this.Ctx.Input.IP()
	if Memcache.IsExist(registerKey) {
		registerStr := Memcache.Get(registerKey).(string)
		log.Println(registerStr)
		registerNumber, err := strconv.Atoi(registerStr)

		if err != nil {
			beego.Error(err)
		}
		if registerNumber > RegisterRecaptchaNumber {
			this.Data["showRecaptcha"] = true
			this.Data["publicKey"] = beego.AppConfig.String("googleRecaptchaPublicKey")
		}
	}

	this.LayoutSections["htmlFooter"] = "footer/registerFooter.html"

	this.Data["xsrfdata"] = template.HTML(this.XsrfFormHtml())

	this.TplNames = "templates/register.html"
}

func (this *RegisterController) Post() {
	userAccount := m.UserAccount{}
	if err := this.ParseForm(&userAccount); err != nil {
		this.Redirect("/r", 302)
	}

	userAccount.Ip = this.Ctx.Input.IP()
	privateKey := beego.AppConfig.String("googleRecaptchaPrivateKey")
	challenge := this.GetString("recaptcha_challenge_field")
	response := this.GetString("recaptcha_response_field")

	var err error
	noNeedRecaptcha := true
	recaptchaCheck := false
	registerNumber := 0
	registerStr := ""
	registerKey := RegisterKeyPrefix + userAccount.Ip
	if Memcache.IsExist(registerKey) {
		registerStr = Memcache.Get(registerKey).(string)
		registerNumber, err = strconv.Atoi(registerStr)
		if err == nil {
			if registerNumber > RegisterRecaptchaNumber {
				noNeedRecaptcha = false
				recaptchaCheck = h.GoogleRecaptcha(privateKey, userAccount.Ip, challenge, response)
			}
		} else {
			beego.Error(err)
		}
	}

	if noNeedRecaptcha || recaptchaCheck {
		userAccount.Salt = h.GetRandomString(5)
		id, err := m.AddUserAccount(&userAccount)
		if err != nil {
			this.Redirect("/r", 302)
		} else {
			if id > 0 {
				if Memcache.IsExist(registerKey) {
					registerNumber++
					registerStr = strconv.Itoa(registerNumber)
					err = Memcache.Put(registerKey, registerStr, RegisterExpired)
					log.Println(err)
				} else {
					Memcache.Put(registerKey, "1", RegisterExpired)
				}

				cookieHash := beego.AppConfig.String("cookieHash")
				cookieName := beego.AppConfig.String("cookieName")
				cookieSep := beego.AppConfig.String("cookieSep")

				cookieStr := strconv.Itoa(int(id)) + cookieSep + userAccount.Email + cookieSep + userAccount.Password
				this.SetSecureCookie(cookieHash, cookieName, cookieStr, 604800)
				this.SetSession("userinfo", userAccount)

				this.Redirect("/", 302)
			}
		}
	} else {
		this.Redirect("/r", 302)
	}
}

func (this *RegisterController) Verify() {

	fieldId := this.GetString("fieldId")
	fieldValue := this.GetString("fieldValue")

	fieldMap := map[string]bool{"nickname": true, "email": true}
	result := [2]interface{}{"nickname", true}
	if _, ok := fieldMap[fieldId]; ok {
		result[0] = fieldId
		count, err := m.GetUserAccountCountByColumn(fieldId, fieldValue)
		if err != nil {
			log.Println(err)
			result[1] = false
		} else {
			if count > 0 {
				result[1] = false
			}
		}
	}

	this.Data["json"] = result
	this.ServeJson()
}
