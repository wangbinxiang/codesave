package controllers

import (
	"codesave/libs"
	m "codesave/models"
	"github.com/astaxie/beego"
)

type IndexController struct {
	libs.BaseController
}

var indexQuestionPageSize int64 = 2

func (this *IndexController) Get() {
	page, _ := this.GetInt(":page")

	if page <= 0 {
		page = 1
	}
	questionIssues, more, err := m.GetQuestionIssueList(page, indexQuestionPageSize)

	if err != nil {
		beego.Error(err)
	} else {
		uids := []int64{}
		for _, v := range questionIssues {
			uids = append(uids, v["Uid"].(int64))
		}

		userAccounts, count, err := m.GetUserAccountListByUids(uids)

		if err != nil {
			beego.Error(err)
		} else {
			userAccountNicknameList := map[int64]string{}
			for _, v := range userAccounts {
				userAccountNicknameList[v["Id"].(int64)] = v["Nickname"].(string)
			}

			for k, v := range questionIssues {
				questionIssues[k]["Nickname"] = userAccountNicknameList[v["Uid"].(int64)]
			}
		}

		if this.IsAjax() {
			this.Data["json"] = map[string]interface{}{"q": questionIssues, "count": count}
			this.ServeJson()
		} else {
			this.Data["q"] = questionIssues
			if more {
				this.Data["next"] = true
				this.Data["nextPage"] = page + 1
			}
			if page > 1 {
				this.Data["prev"] = true
				this.Data["prevPage"] = page - 1
			}
			this.Data["count"] = count
		}

	}

	this.TplNames = "templates/index.html"
}
