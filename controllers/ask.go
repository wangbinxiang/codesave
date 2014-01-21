package controllers

import (
	"github.com/wangbinxiang/codesave/libs"
	m "github.com/wangbinxiang/codesave/models"
	"html/template"
	"log"
	"strconv"
)

type AskController struct {
	libs.BaseController
}

func (this *AskController) Prepare() {
	this.BaseController.Prepare()
	this.LoginJump(true)
}

func (this *AskController) Get() {
	qid, _ := this.GetInt(":qid")

	this.Data["edit"] = false //编辑问题标示
	if qid > 0 {
		questuionIssue, err := m.GetQuestionIssue(qid)

		if err == nil {
			if this.LoginUser.Id != int64(questuionIssue.Uid) {
				this.Redirect("/a", 302)
			}

			this.Data["edit"] = true
			this.Data["q"] = questuionIssue

		} else {
			log.Println(err)
		}
	}

	this.Data["xsrfdata"] = template.HTML(this.XsrfFormHtml())

	this.LayoutSections["htmlFooter"] = "footer/askFooter.html"

	this.TplNames = "templates/ask.html"
}

func (this *AskController) Post() {
	questuionIssue := m.QuestionIssue{}
	if err := this.ParseForm(&questuionIssue); err != nil {
		this.Redirect("/a", 302)
	}
	questuionIssue.Uid = int(this.LoginUser.Id)
	id, err := m.AddQuestionIssue(&questuionIssue)
	if err != nil {
		this.Redirect("/a", 302)
	}

	if id > 0 {
		url := "/q/" + strconv.Itoa(int(id))
		this.Redirect(url, 302)
	}
}

func (this *AskController) Put() {
	questuionIssue := m.QuestionIssue{}
	if err := this.ParseForm(&questuionIssue); err != nil {
		this.Redirect("/a", 302)
	}

	if questuionIssue.Id > 0 {
		_, err := m.GetQuestionIssue(questuionIssue.Id)
		if err != nil {
			log.Println(err)
			this.Redirect("/a", 302)
		} else {
			num, err := m.UpdateQuestionIssue(&questuionIssue)
			log.Println(num, err)
			url := "/a/" + strconv.Itoa(int(questuionIssue.Id))
			this.Redirect(url, 302)
		}
	} else {
		this.Redirect("/a", 302)
	}
}
