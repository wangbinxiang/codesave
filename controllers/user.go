package controllers

import (
	"codesave/libs"
	m "codesave/models"
	"time"
)

type UserController struct {
	libs.BaseController
}

func (this *UserController) Prepare() {
	this.BaseController.Prepare()
	this.LoginJump(true)
}

var userQuestionPageSize int64 = 1

func (this *UserController) Get() {
	page, err := this.GetInt("page")

	if err != nil || page < 2 {
		page = 1
	}

	questionIssues, more, err := m.GetQuestionIssueListByUid(this.LoginUser.Id, page, userQuestionPageSize)

	if this.IsAjax() {
		// fmt.Println(t.Format("20060102 15:04:05"))
		for k, _ := range questionIssues {
			questionIssues[k]["PublishTime"] = questionIssues[k]["PublishTime"].(time.Time).Format("2006-01-02 15:04:05")
			delete(questionIssues[k], "Content")
		}

		result := map[string]interface{}{"q": questionIssues, "more": more}
		this.Data["json"] = result
		this.ServeJson()
	} else {
		this.Data["q"] = questionIssues
		this.Data["more"] = more

		this.LayoutSections["htmlFooter"] = "footer/userFooter.html"
		this.TplNames = "templates/user.html"
	}
}
