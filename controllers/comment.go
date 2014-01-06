package controllers

import (
	"codesave/libs"
	m "codesave/models"
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
				id, err := m.AddCommentInfo(&commentInfo)
				log.Println(err)
				if err == nil {
					result["result"] = true
					result["id"] = id
				}
			}
		}
		this.Data["json"] = result
	}
	this.ServeJson()
}
