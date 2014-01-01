package controllers

import (
	"codesave/libs"
	"codesave/models"
	// "crypto/md5"
	// "encoding/hex"
	"github.com/astaxie/beego"
)

type IndexController struct {
	libs.BaseController
}

func (this *IndexController) Get() {
	user_accounts, err := models.GetAllUserAccount()
	if err != nil {
		beego.Error(err)
	}

	// h := md5.New()
	// h.Write([]byte("testPassword"))

	// username := "testName2"
	// password := hex.EncodeToString(h.Sum(nil))
	// salt := "salts"

	// id, err := models.AddUserAccount(username, password, salt)

	// if err != nil {
	// 	beego.Error(err)
	// }
	// println(id)

	this.Data["user_accounts"] = user_accounts
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "templates/index.html"
}
