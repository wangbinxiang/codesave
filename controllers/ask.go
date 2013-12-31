package controllers

import (
	m "codesave/models"
	"github.com/astaxie/beego"
	"log"
	"strconv"
)

type AskController struct {
	beego.Controller
}

func (this *AskController) Get() {
	qid, _ := this.GetInt(":qid")

	this.Data["edit"] = false //编辑问题标示
	if qid > 0 {
		questuionIssue, err := m.GetQuestionIssue(qid)

		if err == nil {
			this.Data["edit"] = true
			this.Data["q"] = questuionIssue
		} else {
			log.Println(err)
		}
	}

	this.Layout = "layout.html"

	this.TplNames = "templates/ask.html"
}

func (this *AskController) Post() {
	questuionIssue := m.QuestionIssue{}
	if err := this.ParseForm(&questuionIssue); err != nil {
		this.Redirect("/a", 302)
	}

	id, err := m.AddQuestionIssue(&questuionIssue)
	if err != nil {
		this.Redirect("/a", 302)
	}

	if id > 0 {
		url := "/a/" + strconv.Itoa(int(id))
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
