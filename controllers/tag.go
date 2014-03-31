package controllers

import (
	"github.com/wangbinxiang/codesave/libs"
	m "github.com/wangbinxiang/codesave/models"
	"log"
)

type TagController struct {
	libs.BaseController
}

func (this *TagController) Get() {
	tagLabel := m.TagLabel{}
	tagLabel.Id = 1
	tagLabel.Name = "名字2"
	tagLabel.Description = "介绍3"

	id, err := m.UpdateTagLabel(&tagLabel)
	log.Println(id, err)
	var page int64
	page = 1
	log.Println(123)
	tagInfos, more, count, err := m.GetTagLabelList(page, 20, false)
	log.Println(tagInfos, more, count, err)

	this.Redirect("/", 302)
}
