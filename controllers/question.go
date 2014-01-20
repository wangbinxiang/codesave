package controllers

import (
	"codesave/libs"
	m "codesave/models"
	"github.com/astaxie/beego"
	"log"
)

type QuestionController struct {
	libs.BaseController
}

var commentPageSize int64 = 10

func (this *QuestionController) Get() {
	qid, _ := this.GetInt(":qid")

	if qid > 0 {
		questuionIssue, err := m.GetQuestionIssue(qid)
		if err == nil {
			//获取评论
			commentInfos, more, err := m.GetCommentInfoListByQid(questuionIssue.Id, 1, commentPageSize)
			if err == nil {
				if len(commentInfos) > 0 {
					uids := []int64{}
					for _, v := range commentInfos {
						uids = append(uids, v["Uid"].(int64))
					}

					userAccounts, _, err := m.GetUserAccountListByUids(uids)

					if err != nil {
						beego.Error(err)
					} else {
						userAccountNicknameList := map[int64]string{}
						for _, v := range userAccounts {
							userAccountNicknameList[v["Id"].(int64)] = v["Nickname"].(string)
						}

						for k, v := range commentInfos {
							commentInfos[k]["Nickname"] = userAccountNicknameList[v["Uid"].(int64)]
						}
					}
					this.Data["cMore"] = more
					this.Data["c"] = commentInfos
				}
			}

			this.Data["q"] = questuionIssue
		} else {
			log.Println(err)
			this.Redirect("/", 302)
		}

	} else {
		this.Ctx.Redirect(302, "/")
	}

	this.LayoutSections["htmlFooter"] = "footer/questionFooter.html"

	this.TplNames = "templates/question.html"
}

func (this *QuestionController) GetComment() {
	qid, _ := this.GetInt("qid")
	page, _ := this.GetInt("page")

	if qid > 0 {
		if page < 2 {
			page = 2
		}

		commentInfos, more, err := m.GetCommentInfoListByQid(qid, page, commentPageSize)
		if err == nil {

			uids := []int64{}
			for _, v := range commentInfos {
				uids = append(uids, v["Uid"].(int64))
			}

			userAccounts, _, err := m.GetUserAccountListByUids(uids)

			if err != nil {
				beego.Error(err)
			} else {
				userAccountNicknameList := map[int64]string{}
				for _, v := range userAccounts {
					userAccountNicknameList[v["Id"].(int64)] = v["Nickname"].(string)
				}

				for k, v := range commentInfos {
					commentInfos[k]["Nickname"] = userAccountNicknameList[v["Uid"].(int64)]
				}
			}

			this.Data["json"] = map[string]interface{}{"c": commentInfos, "cMore": more}
		}
	}
	this.ServeJson()
}
