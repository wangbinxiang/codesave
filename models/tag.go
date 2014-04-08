package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"time"
)

type TagLabel struct {
	Id          int64
	Name        string `orm:"size(50)" "valid:"MinSize(1);MaxSize(50)"`
	Description string `orm:"size(255)" "valid:"MaxSize(255)"`
	FollowNum   uint
	QuestionNum uint
	Status      int8      `orm:"Min(0);Max(1)"`
	CreateTime  time.Time `orm:"index;auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`
}

func (t *TagLabel) TableEngine() string {
	return "INNODB"
}

func checkTagLabel(t *TagLabel) error {
	valid := validation.Validation{}
	b, _ := valid.Valid(t)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func init() {
	MysqlRegisterModelWithPrefix(new(TagLabel))
}

func AddTagLabel(t *TagLabel) (int64, error) {
	if err := checkTagLabel(t); err != nil {
		return 0, err
	}

	id, err := Orm.Insert(t)

	return id, err
}

func UpdateTagLabel(t *TagLabel) (int64, error) {
	if err := checkTagLabel(t); err != nil {
		return 0, err
	}

	node := make(orm.Params)

	if len(t.Name) > 0 {
		node["Name"] = t.Name
	}
	if len(t.Description) > 0 {
		node["Description"] = t.Description
	}
	if t.Status != 0 {
		node["Status"] = t.Status
	}
	if len(node) == 0 {
		return 0, errors.New("update field is empty")
	}

	var table TagLabel
	num, err := Orm.QueryTable(table).Filter("Id", t.Id).Update(node)
	return num, err
}

func GetTagLabelList(page int64, pageSize int64, getTotal bool) ([]orm.Params, bool, int64, error) {
	var (
		tagLabels []orm.Params
		offset    int64
		table     TagLabel
		more      bool
		count     int64
		err       error
	)

	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}
	if getTotal {
		_, err = Orm.QueryTable(table).Limit(pageSize, offset).OrderBy("-id").Values(&tagLabels)
		count, err = Orm.QueryTable(table).Count()
	} else {
		count, err = Orm.QueryTable(table).Limit(pageSize+1, offset).OrderBy("-id").Values(&tagLabels)
		if count > pageSize {
			more = true
			tagLabels = tagLabels[:pageSize]
		}
	}

	return tagLabels, more, count, err
}

func GetAllTagLabelList() ([]orm.Params, int64, error) {
	var (
		tagLabels []orm.Params
		table     TagLabel
		count     int64
		err       error
	)

	count, err = Orm.QueryTable(table).OrderBy("-id").Values(&tagLabels)
	return tagLabels, count, err
}

func GetTagLabelListByIds(ids []int64) ([]orm.Params, int64, error) {
	var (
		tagLabels []orm.Params
		table     TagLabel
		count     int64
		err       error
	)
	count, err = Orm.QueryTable(table).Filter("id__in", ids).Values(&tagLabels)
	return tagLabels, count, err
}
