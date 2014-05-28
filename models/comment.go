package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"time"
)

type CommentInfo struct {
	Id              int64
	Content         string         `orm:"size(255)" valid:"MinSize(5);MaxSize(255)"`
	Left            string         `orm:"size(255)" valid:"MaxSize(255);Match(/^[0-9]\\d*\\.?\\d*$/)"`
	Top             string         `orm:"size(255)" valid:"MaxSize(255);Match(/^[0-9]\\d*\\.?\\d*$/)"`
	PublishTime     time.Time      `orm:"auto_now_add;type(datetime)"`
	QuestionIssue   *QuestionIssue `orm:"rel(fk)"`
	User_account_id int64          `orm:"index"`
}

func (c *CommentInfo) TableIndex() [][]string {
	return [][]string{
		[]string{"question_issue_id", "Id"},
	}
}

func (c *CommentInfo) TableEngine() string {
	return "INNODB"
}

func checkCommentInfo(c *CommentInfo) error {
	valid := validation.Validation{}
	b, _ := valid.Valid(c)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func init() {
	MysqlRegisterModelWithPrefix(new(CommentInfo))
}

func AddCommentInfo(c *CommentInfo) (int64, error) {
	if err := checkCommentInfo(c); err != nil {
		return 0, err
	}
	id, err := Orm.Insert(c)
	return id, err
}

func GetCommentInfoListByQid(qid int64, page int64, pageSize int64) ([]orm.Params, bool, error) {
	var (
		commentInfos []orm.Params
		offset       int64
		table        CommentInfo
	)
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}

	count, err := Orm.QueryTable(table).Filter("question_issue_id", qid).Limit(pageSize+1, offset).OrderBy("-id").Values(&commentInfos)

	more := false
	if count > pageSize {
		more = true
		commentInfos = commentInfos[:pageSize]
	}

	return commentInfos, more, err
}
