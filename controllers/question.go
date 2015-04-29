package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wangbinxiang/codesave/libs"
	m "github.com/wangbinxiang/codesave/models"
	"log"
)

type QuestionController struct {
	libs.BaseController
}

const CommentPageSize int64 = 10

func (this *QuestionController) Get() {
	qid, _ := this.GetInt(":qid")

	if qid > 0 {
		questuionIssue, err := m.GetQuestionIssue(int64(qid))
		if err == nil {
			//获取评论
			commentInfos, more, err := m.GetCommentInfoListByQid(questuionIssue.Id, 1, CommentPageSize)
			if err == nil {
				if len(commentInfos) > 0 {
					questionTags, count, err := m.GetQuestionTagListByQid(questuionIssue.Id)
					log.Println(questionTags, count, err)
					uids := []int64{}
					for _, v := range commentInfos {
						if v["User_account_id"].(int64) > 0 {
							uids = append(uids, v["User_account_id"].(int64))
						}
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
							if v["User_account_id"].(int64) > 0 {
								commentInfos[k]["Nickname"] = userAccountNicknameList[v["User_account_id"].(int64)]
							} else {
								commentInfos[k]["Nickname"] = "游客"
							}
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

		commentInfos, more, err := m.GetCommentInfoListByQid(int64(qid), int64(page), CommentPageSize)
		if err == nil {

			uids := []int64{}
			for _, v := range commentInfos {
				uids = append(uids, v["UserAccount"].(int64))
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
					commentInfos[k]["Nickname"] = userAccountNicknameList[v["UserAccount"].(int64)]
				}
			}

			this.Data["json"] = map[string]interface{}{"c": commentInfos, "cMore": more}
		}
	}
	this.ServeJson()
}
