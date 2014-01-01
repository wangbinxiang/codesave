package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"time"
)

type QuestionIssue struct {
	Id          int64
	Title       string    `orm:"size(32)"  valid:"Required;MinSize(5);MaxSize(150)"`
	Content     string    `orm:"type(text)"  valid:"Required;MinSize(10);MaxSize(30000)"`
	Uid         int64     `orm:"index"`
	PublishTime time.Time `orm:"index;auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`
	CommentNum  uint
}

// 设置引擎为 INNODB
func (q *QuestionIssue) TableEngine() string {
	return "INNODB"
}

//验证信息
func checkQuestionIssue(q *QuestionIssue) error {
	valid := validation.Validation{}
	b, _ := valid.Valid(q)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func init() {
	MysqlRegisterModelWithPrefix(new(QuestionIssue))
}

func AddQuestionIssue(q *QuestionIssue) (int64, error) {
	if err := checkQuestionIssue(q); err != nil {
		return 0, err
	}

	id, err := Orm.Insert(q)

	return id, err
}

func GetQuestionIssue(qid int64) (QuestionIssue, error) {
	questionIssue := QuestionIssue{Id: qid}

	err := Orm.Read(&questionIssue)

	return questionIssue, err
}

func UpdateQuestionIssue(q *QuestionIssue) (int64, error) {
	if err := checkQuestionIssue(q); err != nil {
		return 0, err
	}
	questionIssue := make(orm.Params)
	questionIssue["Title"] = q.Title
	questionIssue["Content"] = q.Content
	var table QuestionIssue
	num, err := Orm.QueryTable(table).Filter("Id", q.Id).Update(questionIssue)
	return num, err
}
