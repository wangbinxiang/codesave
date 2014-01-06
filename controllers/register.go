package controllers

import (
	h "codesave/helper"
	"codesave/libs"
	m "codesave/models"
	"log"
)

type RegisterController struct {
	libs.BaseController
}

func (this *RegisterController) Prepare() {
	this.BaseController.Prepare()
	this.LoginJump(false)
}

func (this *RegisterController) Get() {

	this.LayoutSections["htmlFooter"] = "footer/registerFooter.html"

	this.TplNames = "templates/register.html"
}

func (this *RegisterController) Post() {
	userAccount := m.UserAccount{}
	if err := this.ParseForm(&userAccount); err != nil {
		this.Redirect("/r", 302)
	}

	userAccount.Ip = this.Ctx.Input.IP()

	userAccount.Salt = h.GetRandomString(5)

	id, err := m.AddUserAccount(&userAccount)
	if err != nil {
		this.Redirect("/r", 302)
	}

	if id > 0 {
		this.Redirect("/", 302)
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
