package controllers

import (
	"github.com/wangbinxiang/codesave/libs"
	m "github.com/wangbinxiang/codesave/models"
	"log"
)

type CommentController struct {
	libs.BaseController
}

func (this *CommentController) Post() {
	if this.IsAjax() {
		result := map[string]interface{}{"result": false, "id": 0}
		qid, _ := this.GetInt("Qid")
		if qid > 0 {
			commentInfo := m.CommentInfo{}
			if err := this.ParseForm(&commentInfo); err != nil {
				log.Println(err)
			} else {
				commentInfo.QuestionIssue = new(m.QuestionIssue)
				commentInfo.QuestionIssue.Id = qid
				_, err := m.GetQuestionIssue(int64(commentInfo.QuestionIssue.Id))
				if err == nil {
					// commentInfo.UserAccount = new(m.UserAccount)
					if this.IsLogin {
						commentInfo.User_account_id = int64(this.LoginUser.Id)
					}

					err := m.Orm.Begin()
					log.Println(commentInfo)
					id, err := m.AddCommentInfo(&commentInfo)
					if err != nil || id == 0 {
						m.Orm.Rollback()
					} else {
						num, err := m.AddQuestionIssueCommentNum(int64(commentInfo.QuestionIssue.Id))
						if err != nil || num == 0 {
							m.Orm.Rollback()
						} else {
							m.Orm.Commit()

							result["result"] = true
							result["id"] = id
						}
					}
				}
			}
		}
		this.Data["json"] = result
	}
	this.ServeJson()
}
