package controllers

import (
	m "codesave/models"
	// "crypto/md5"
	// "encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
)

type AskController struct {
	beego.Controller
}

func (this *AskController) Get() {
	this.Layout = "layout.html"

	this.TplNames = "templates/ask.html"
}

func (this *AskController) Post() {
	questuionIssue := m.QuestionIssue{}
	fmt.Print(333)
	if err := this.ParseForm(&questuionIssue); err != nil {
		this.Ctx.Redirect(302, "/")
	}
	fmt.Print(questuionIssue)
	qid, _ := this.GetInt("id")

	if qid > 0 {
		fmt.Print(qid)

		this.Data["Intid"] = qid
	} else {
		fmt.Print(1233)
		id, err := m.AddQuestionIssue(&questuionIssue)
		if err != nil {
			beego.Error(err)
		}
		this.Data["Intid"] = id
	}

	this.Layout = "layout.html"

	this.TplNames = "templates/ask.html"
}
