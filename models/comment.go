package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"time"
)

type CommentInfo struct {
	Id          int64
	Qid         int       `orm:"index" valid:"Max(100);Min(1)"`
	Uid         int       `orm:"index" valid:"Max(100);Min(1)"`
	Content     string    `orm:"size(255)" valid:"MinSize(5);MaxSize(765)"`
	Left        string    `orm:"size(255)" valid:"MaxSize(255);Match(/^[0-9]\\d*\\.?\\d*$/)"`
	Top         string    `orm:"size(255)" valid:"MaxSize(255);Match(/^[0-9]\\d*\\.?\\d*$/)"`
	PublishTime time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c *CommentInfo) TableIndex() [][]string {
	return [][]string{
		[]string{"Qid", "Id"},
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

func GetCommentInfoListByQid(qid int64, page int64, page_size int64) ([]orm.Params, int64, error) {
	var commentInfos []orm.Params
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}

	var table CommentInfo
	count, err := Orm.QueryTable(table).Filter("qid", qid).Limit(page_size, offset).OrderBy("-id").Values(&commentInfos)

	return commentInfos, count, err
}
