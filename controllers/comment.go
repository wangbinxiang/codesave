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
	if !this.IsLogin {
		if this.IsAjax() {
			this.StopRun()
		} else {
			this.Redirect("/", 302)
		}
	}
}

func (this *CommentController) Post() {
	if this.IsAjax() {
		commentInfo := m.CommentInfo{}
		result := map[string]interface{}{"result": false, "id": 0}
		if err := this.ParseForm(&commentInfo); err != nil {
			log.Println(err)
		} else {
			log.Println(&commentInfo)
			id, err := m.AddCommentInfo(&commentInfo)
			if err == nil {
				result["result"] = true
				result["id"] = id
			}
		}
		this.Data["json"] = result
	}
	this.ServeJson()
}
