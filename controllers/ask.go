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

	//get all tag info
	tagLabels, tagCount, err := m.GetAllTagLabelList()

	if err != nil {
		log.Println(err)
	} else {
		this.Data["tagCount"] = tagCount
		log.Println(tagCount)
		this.Data["tagLabels"] = tagLabels
	}

	if qid > 0 {
		questuionIssue, err := m.GetQuestionIssue(qid)

		if err == nil {
			if this.LoginUser.Id != int64(questuionIssue.UserAccount.Id) {
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
	questuionIssue.UserAccount = new(m.UserAccount)
	questuionIssue.UserAccount.Id = int64(this.LoginUser.Id)
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
			url := "/q/" + strconv.Itoa(int(questuionIssue.Id))
			this.Redirect(url, 302)
		}
	} else {
		this.Redirect("/a", 302)
	}
}
