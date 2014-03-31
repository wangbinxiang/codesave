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
	Title       string    `orm:"size(32)"  valid:"MinSize(5);MaxSize(32)"`
	Content     string    `orm:"type(text)"  valid:"MinSize(10)"`
	PublishTime time.Time `orm:"index;auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`
	CommentNum  uint
	UserAccount *UserAccount   `orm:"rel(fk)"`
	CommentInfo []*CommentInfo `orm:"reverse(many)"`
	QuestionTag []*QuestionTag `orm:"reverse(many)"`
}

type QuestionTag struct {
	Id            int64
	TagLabel      *TagLabel      `orm:"rel(fk)"`
	QuestionIssue *QuestionIssue `orm:"rel(fk)"`
	CreateTime    time.Time      `orm:auto_now_add;type(datetime)`
}

// 设置引擎为 INNODB
func (q *QuestionIssue) TableEngine() string {
	return "INNODB"
}

func (q *QuestionTag) TableEngine() string {
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

func checkQuestionTag(q *QuestionTag) error {
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
	MysqlRegisterModelWithPrefix(new(QuestionIssue), new(QuestionTag))
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

func GetQuestionIssueList(page int64, pageSize int64) ([]orm.Params, bool, error) {
	var (
		questionIssues []orm.Params
		offset         int64
		table          QuestionIssue
	)
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}

	count, err := Orm.QueryTable(table).Limit(pageSize+1, offset).OrderBy("-id").Values(&questionIssues)
	more := false
	if count > pageSize {
		more = true
		questionIssues = questionIssues[:pageSize]
	}

	return questionIssues, more, err
}

func GetQuestionIssueListByUid(uid int64, page int64, pageSize int64) ([]orm.Params, bool, error) {
	var questionIssues []orm.Params
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}

	var table QuestionIssue
	count, err := Orm.QueryTable(table).Limit(pageSize+1, offset).Filter("user_account_id", uid).OrderBy("-id").Values(&questionIssues)
	more := false
	if count > pageSize {
		more = true
		questionIssues = questionIssues[:pageSize]
	}

	return questionIssues, more, err

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

func AddQuestionIssueCommentNum(qid int64) (int64, error) {
	var table QuestionIssue
	num, err := Orm.QueryTable(table).Filter("Id", qid).Update(orm.Params{"comment_num": orm.ColValue(orm.Col_Add, 1)})
	return num, err
}

func AddQUestionTag(q *QuestionTag) (int64, error) {
	if err := checkQuestionTag(q); err != nil {
		return 0, err
	}

	id, err := Orm.Insert(q)

	return id, err
}

func GetQuestionTagListByQid(qid int64) ([]orm.Params, int64, error) {
	var (
		questionTags []orm.Params
		table        QuestionTag
		count        int64
		err          error
	)
	count, err = Orm.QueryTable(table).Filter("question_issue_id", qid).Values(&questionTags)
	return questionTags, count, err
}
