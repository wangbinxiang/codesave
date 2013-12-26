package controllers

import (
	// "codesave/models"
	// "crypto/md5"
	// "encoding/hex"
	"github.com/astaxie/beego"
)

type QuestionController struct {
	beego.Controller
}

func (this *QuestionController) Get() {
	this.Layout = "layout.html"

	this.TplNames = "templates/question.html"
}

func (this *QuestionController) Post() {

}
