package controllers

import (
	"codesave/libs"
	m "codesave/models"
	"log"
)

type QuestionController struct {
	libs.BaseController
}

func (this *QuestionController) Get() {
	qid, _ := this.GetInt(":qid")

	if qid > 0 {
		questuionIssue, err := m.GetQuestionIssue(qid)
		if err == nil {
			//获取评论
			commentInfos, count, _ := m.GetCommentInfoListByQid(questuionIssue.Id, 1, 20)
			if count > 0 {
				this.Data["c"] = commentInfos
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

		commentInfos, count, _ := m.GetCommentInfoListByQid(qid, page, 20)
		if count > 0 {
			this.Data["json"] = map[string]interface{}{"c": commentInfos}
		}
	}
	this.ServeJson()
}
