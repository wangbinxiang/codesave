package controllers

import (
	// "codesave/models"
	// "crypto/md5"
	// "encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type AskController struct {
	beego.Controller
}

func (this *AskController) Get() {
	this.Layout = "layout.html"

	this.TplNames = "templates/ask.html"
}

func (this *AskController) Post() {
	this.Layout = "layout.html"

	this.TplNames = "templates/ask.html"

	id := this.Input().Get("id")
	intid, err := strconv.Atoi(id)

	fmt.Println(err)
	fmt.Println(intid)

	this.Data["Intid"] = intid
}
