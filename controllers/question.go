package controllers

import (
	"codesave/libs"
	m "codesave/models"
	// "crypto/md5"
	// "encoding/hex"
	"log"
)

type QuestionController struct {
	libs.BaseController
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

	this.LayoutSections["htmlHeader"] = "header/askHeader.html"

	this.TplNames = "templates/question.html"
}
