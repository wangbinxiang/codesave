package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"github.com/wangbinxiang/codesave/helper"
	"github.com/wangbinxiang/codesave/libs"
	m "github.com/wangbinxiang/codesave/models"
	"log"
)

type TagController struct {
	libs.BaseController
}

const TagPageSize int64 = 2

func (this *TagController) Get() {
	name := this.GetString(":name")

	if len(name) > 0 {
		page, err := this.GetInt("p")
		if err != nil {
			page = 1
		}
		tagLabel := m.TagLabel{Name: name}
		num, err := m.GetTabLabelByName(&tagLabel, int64(page), TagPageSize)
		if err != nil {
			log.Println(num, err)
			this.Redirect("/t", 302)
		}
		this.Data["tagLabel"] = tagLabel
		if num > 0 {
			qids := make([]int64, 0, num)
			for _, v := range tagLabel.QuestionTags {
				qids = append(qids, v.QuestionIssue.Id)
			}
			log.Println(qids)
			questionIssues, qNum, err := m.GetQuestionIssueListByIds(qids)
			if err != nil {
				log.Println(err)
				this.Redirect("/t", 302)
			}
			if qNum > 0 {
				uids := make([]int64, 0, qNum)
				for _, v := range questionIssues {
					uids = append(uids, v.UserAccount.Id)
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
					this.Data["userAccountNicknameList"] = userAccountNicknameList
				}

				this.Data["qNum"] = qNum
				this.Data["questionIssues"] = questionIssues
			}
		}
	} else {

	}

	this.TplNames = "templates/tag.html"
}
