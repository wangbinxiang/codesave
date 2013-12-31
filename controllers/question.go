package controllers

import (
	m "codesave/models"
	// "crypto/md5"
	// "encoding/hex"
	"github.com/astaxie/beego"
	"log"
)

type QuestionController struct {
	beego.Controller
}

func (this *QuestionController) Get() {
	qid, _ := this.GetInt(":qid")

	if qid > 0 {
		questuionIssue, err := m.GetQuestionIssue(qid)
		if err == nil {
			this.Data["q"] = questuionIssue
		} else {
			log.Println(err)
		}
	} else {
		this.Ctx.Redirect(302, "/")
	}

	this.Layout = "layout.html"

	this.TplNames = "templates/question.html"
}
