package models

import (
	"codesave/libs"
	"time"
)

type QuestionIssue struct {
	Id          int64
	Title       string    `orm:"size(32)"`
	Content     string    `orm:"type(text)"`
	Uid         int64     `orm:"index"`
	PublishTime time.Time `orm:"index;auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`
	CommentNum  int64
}

// 设置引擎为 INNODB
func (q *QuestionIssue) TableEngine() string {
	return "INNODB"
}

func init() {
	libs.MysqlRegisterModelWithPrefix(new(QuestionIssue))
}

func AddQuestionIssue(q *QuestionIssue) (int64, error) {
	id, err := libs.Orm.Insert(q)

	return id, err
}

func GetQuestionIssue(qid int64) (QuestionIssue, error) {
	questionIssue := QuestionIssue{Id: qid}

	err := libs.Orm.Read(&questionIssue)

	return questionIssue, err
}
