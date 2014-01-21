package controllers

import (
	"github.com/wangbinxiang/codesave/libs"
	m "github.com/wangbinxiang/codesave/models"
	"log"
)

type CommentController struct {
	libs.BaseController
}

func (this *CommentController) Prepare() {
	this.BaseController.Prepare()
	this.LoginJump(true)
}

func (this *CommentController) Post() {
	if this.IsAjax() {
		commentInfo := m.CommentInfo{}
		result := map[string]interface{}{"result": false, "id": 0}
		if err := this.ParseForm(&commentInfo); err != nil {
			log.Println(err)
		} else {
			_, err := m.GetQuestionIssue(int64(commentInfo.Qid))
			if err == nil {
				commentInfo.Uid = int(this.LoginUser.Id)

				err := m.Orm.Begin()

				id, err := m.AddCommentInfo(&commentInfo)
				if err != nil || id == 0 {
					m.Orm.Rollback()
				} else {
					num, err := m.AddQuestionIssueCommentNum(int64(commentInfo.Qid))
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
		this.Data["json"] = result
	}
	this.ServeJson()
}
