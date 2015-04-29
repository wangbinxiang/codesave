package controllers

import (
	"github.com/wangbinxiang/codesave/helper"
	"github.com/wangbinxiang/codesave/libs"
	m "github.com/wangbinxiang/codesave/models"
	"html/template"
	"log"
	"strconv"
)

type AskController struct {
	libs.BaseController
}

const (
	tagsMinLen int64 = 0
	tagsMaxLen int64 = 5
)

func (this *AskController) Prepare() {
	this.BaseController.Prepare()
	this.LoginJump(true)
}

func (this *AskController) Get() {
	qid, _ := this.GetInt(":qid")

	this.Data["edit"] = false //编辑问题标示

	//get all tag info
	tagLabels, tagCount, err := m.GetAllEnableTagLabelList()

	if err != nil {
		log.Println(err)
	} else {
		this.Data["tagCount"] = tagCount
		this.Data["tagLabels"] = tagLabels
	}

	if qid > 0 {
		questuionIssue, err := m.GetQuestionIssue(int64(qid))

		if err == nil {
			if this.LoginUser.Id != int64(questuionIssue.UserAccount.Id) {
				this.Redirect("/a", 302)
			}

			questionTags, count, err := m.GetQuestionTagListByQid(questuionIssue.Id)
			if err != nil {
				log.Println(err)
				this.Redirect("/a", 302)
			}
			questionTagInfos := make([]interface{}, 0, count)
			for _, v := range questionTags {
				for _, vv := range tagLabels {
					if vv["Id"] == v["TagLabel"] {
						questionTagInfos = append(questionTagInfos, vv)
					}
				}
			}

			this.Data["edit"] = true
			this.Data["q"] = questuionIssue
			this.Data["qt"] = questionTagInfos
		} else {
			log.Println(err)
			this.Redirect("/a", 302)
		}
	}

	this.Data["xsrfdata"] = template.HTML(this.XsrfFormHtml())

	this.LayoutSections["htmlFooter"] = "footer/askFooter.html"

	this.TplNames = "templates/ask.html"
}

func (this *AskController) Post() {
	questuionIssue := m.QuestionIssue{}
	if err := this.ParseForm(&questuionIssue); err != nil {
		this.Redirect("/a", 302)
	}

	tags := helper.SliceStringToInt64(this.GetStrings("tags"))
	checkTagRes := checkTags(tags)
	if !checkTagRes {
		this.Redirect("/a", 302)
	}

	m.Orm.Begin()

	questuionIssue.UserAccount = new(m.UserAccount)
	questuionIssue.UserAccount.Id = int64(this.LoginUser.Id)
	id, err := m.AddQuestionIssue(&questuionIssue)
	if err != nil {
		m.Orm.Rollback()
		this.Redirect("/a", 302)
	}
	_, err = m.AddQuestionTagMulti(&questuionIssue, tags)
	if err != nil {
		m.Orm.Rollback()
		this.Redirect("/a", 302)
	}

	_, err = m.AddTagLabelQuestionNum(tags)
	if err != nil {
		m.Orm.Rollback()
		this.Redirect("/a", 302)
	}

	if id > 0 {
		m.Orm.Commit()
		url := "/q/" + strconv.Itoa(int(id))
		this.Redirect(url, 302)
	}
}

func (this *AskController) Put() {
	questuionIssue := m.QuestionIssue{}
	if err := this.ParseForm(&questuionIssue); err != nil {
		this.Redirect("/a", 302)
	}
	tags := helper.SliceStringToInt64(this.GetStrings("tags"))
	checkTagRes := checkTags(tags)
	log.Println(checkTagRes)
	if !checkTagRes {
		this.Redirect("/a", 302)
	}

	if questuionIssue.Id > 0 {
		oldQuestionIssue, err := m.GetQuestionIssue(questuionIssue.Id)
		if err != nil || oldQuestionIssue.UserAccount.Id != int64(this.LoginUser.Id) {
			this.Redirect("/a", 302)
		} else {
			questionTags, count, err := m.GetQuestionTagListByQid(questuionIssue.Id)
			questionTagIds := make(map[int64]bool, count)
			newTags := make(map[int64]bool, len(tags))

			if count > 0 {
				for _, v := range questionTags {
					questionTagIds[v["TagLabel"].(int64)] = true
				}
				for _, v := range tags {
					newTags[v] = true
					if _, ok := questionTagIds[v]; ok {
						delete(questionTagIds, v)
						delete(newTags, v)
					}
				}
			}

			m.Orm.Begin()

			num, err := m.UpdateQuestionIssue(&questuionIssue)
			if err != nil {
				m.Orm.Rollback()
			}
			log.Println(num, err)
			log.Println("questionTagIds:", questionTagIds)
			delTagsLen := len(questionTagIds)
			if delTagsLen > 0 {
				delTagIds := make([]int64, 0, delTagsLen)
				for _, v := range questionTags {
					if questionTagIds[v["TagLabel"].(int64)] {
						delTagIds = append(delTagIds, v["TagLabel"].(int64))
						num, err = m.DelQuestionTag(&m.QuestionTag{Id: v["Id"].(int64)})
						log.Println(num, err)
						if err != nil {
							m.Orm.Rollback()
						}
					}
				}

				_, err = m.MinusTagLabelQuestionNum(delTagIds)
				if err != nil {
					m.Orm.Rollback()
					this.Redirect("/a", 302)
				}
			}

			if len(newTags) > 0 {
				tags = make([]int64, 0, len(newTags))
				for k, _ := range newTags {
					tags = append(tags, k)
				}
				_, err = m.AddQuestionTagMulti(&questuionIssue, tags)
				if err != nil {
					m.Orm.Rollback()
				}

				_, err = m.AddTagLabelQuestionNum(tags)
				if err != nil {
					m.Orm.Rollback()
					this.Redirect("/a", 302)
				}
			}

			m.Orm.Commit()
			url := "/q/" + strconv.Itoa(int(questuionIssue.Id))
			this.Redirect(url, 302)
		}
	} else {
		this.Redirect("/a", 302)
	}
}

func checkTags(tags []int64) bool {
	tagsLen := int64(len(tags))
	if tagsLen == 0 || tagsLen > 5 {
		return false
	}

	_, count, _ := m.GetTagLabelListByIds(tags)
	if count == 0 {
		return false
	}
	if count != tagsLen {
		return false
	}
	return true
}
