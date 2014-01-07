package models

import (
	h "codesave/helper"
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"time"
)

type UserAccount struct {
	Id           int64
	Email        string    `orm:"size(128),unique" valid:"Email;MaxSize(128)"`
	Nickname     string    `orm:"size(32),unique" valid:"MinSize(2);MaxSize(32);Match(/^([^\\x00-\\xff\\s]|[0-9a-zA-Z_])+$/)"`
	Password     string    `orm:"size(32)" valid:"MinSize(6);MaxSize(16)"`
	Salt         string    `orm:"size(5)" valid:"Length(5)"`
	RegisterTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime   time.Time `orm:"auto_now;type(datetime)"`
	Ip           string    `orm:"size(32)" valid:"IP"`
}

// 设置引擎为 INNODB
func (u *UserAccount) TableEngine() string {
	return "INNODB"
}

func checkUserAccount(u *UserAccount) error {
	valid := validation.Validation{}
	b, _ := valid.Valid(u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func init() {
	MysqlRegisterModelWithPrefix(new(UserAccount))
}

func AddUserAccount(u *UserAccount) (int64, error) {
	if err := checkUserAccount(u); err != nil {
		return 0, err
	}

	u.Password = h.MD5(u.Password + u.Salt)

	id, err := Orm.Insert(u)

	return id, err
}

func GetAllUserAccount() ([]*UserAccount, error) {
	userAccount := make([]*UserAccount, 0)

	qs := Orm.QueryTable(new(UserAccount))

	_, err := qs.All(&userAccount)
	return userAccount, err
}

func GetUserAccountListByUids(uids []int64) ([]orm.Params, int64, error) {
	var userAccounts []orm.Params
	var table UserAccount
	count, err := Orm.QueryTable(table).Filter("id__in", uids).Values(&userAccounts)

	return userAccounts, count, err
}

func GetUserAccountCountByColumn(column string, columnValue string) (int64, error) {
	var table UserAccount

	count, err := Orm.QueryTable(table).Filter(column, columnValue).Count()

	return count, err
}

func GetUserAccountByEmail(email string) (UserAccount, error) {
	var table UserAccount
	userAccount := UserAccount{Email: email}

	err := Orm.QueryTable(table).Filter("email", email).One(&userAccount)

	return userAccount, err
}
