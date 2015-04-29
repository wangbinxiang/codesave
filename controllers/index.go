package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"github.com/wangbinxiang/codesave/helper"
	"github.com/wangbinxiang/codesave/libs"
	m "github.com/wangbinxiang/codesave/models"
	"log"
)

type IndexController struct {
	libs.BaseController
}

const indexQuestionPageSize int64 = 20

func (this *IndexController) Get() {
	page, _ := this.GetInt(":page")

	if page <= 0 {
		page = 1
	}
	questionIssues, more, err := m.GetQuestionIssueList(int64(page), indexQuestionPageSize)

	if err != nil {
		beego.Error(err)
	} else {

		if len(questionIssues) > 0 {
			uids := []int64{}
			for _, v := range questionIssues {
				uids = append(uids, v["UserAccount"].(int64))
			}

			uidSliceInterface, ok := helper.TakeSliceArg(uids)
			if !ok {
				log.Println("uids error")
				this.Redirect("/t", 302)
			}
			uidSliceInterface = utils.SliceUnique(uidSliceInterface)
			uids = helper.SliceInterfaceConvert(uidSliceInterface).([]int64)
			userAccounts, _, err := m.GetUserAccountListByUids(uids)

			if err != nil {
				beego.Error(err)
			} else {
				userAccountNicknameList := map[int64]string{}
				for _, v := range userAccounts {
					userAccountNicknameList[v["Id"].(int64)] = v["Nickname"].(string)
				}

				for k, v := range questionIssues {
					questionIssues[k]["Nickname"] = userAccountNicknameList[v["UserAccount"].(int64)]
				}
			}
		}

		if this.IsAjax() {
			this.Data["json"] = map[string]interface{}{"q": questionIssues, "more": more}
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
			this.Data["more"] = more
		}

	}

	this.TplNames = "templates/index.html"
}
