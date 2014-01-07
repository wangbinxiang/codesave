package controllers

import (
	"codesave/libs"
	m "codesave/models"
	"github.com/astaxie/beego"
	"log"
)

type IndexController struct {
	libs.BaseController
}

func (this *IndexController) Get() {

	page, _ := this.GetInt("page")

	if page <= 0 {
		page = 1
	}
	var pageSize int64
	pageSize = 20

	questionIssues, count, err := m.GetQuestionIssueList(page, pageSize)
	log.Println(count)
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
			userAccountList := map[int64]interface{}{}
			for _, v := range userAccounts {
				log.Println(v)
				userAccountList[v["Id"].(int64)] = v
			}
			this.Data["u"] = userAccountList
		}

		this.Data["q"] = questionIssues
		this.Data["count"] = count
	}

	this.TplNames = "templates/index.html"
}
