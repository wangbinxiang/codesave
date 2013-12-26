package models

import (
	"codesave/libs"
	"github.com/astaxie/beego/orm"
	"time"
)

type UserAccount struct {
	Id           int64
	Username     string    `orm:"size(32),unique"`
	Password     string    `orm:"size(32)"`
	Salt         string    `orm:"size(5)"`
	RegisterTime time.Time `orm:"index;auto_now_add;type(datetime)"`
	UpdateTime   time.Time `orm:"auto_now;type(datetime)"`
}

// 设置引擎为 INNODB
func (u *UserAccount) TableEngine() string {
	return "INNODB"
}

func init() {
	libs.MysqlRegisterModelWithPrefix(new(UserAccount))
}

func AddUserAccount(username, password, salt string) (id int64, err error) {
	o := orm.NewOrm()

	register_time := time.Now()
	update_time := time.Now()

	userAccount := &UserAccount{Username: username, Password: password, Salt: salt, RegisterTime: register_time, UpdateTime: update_time}

	id, err = o.Insert(userAccount)

	return id, err
}

func GetAllUserAccount() ([]*UserAccount, error) {
	o := orm.NewOrm()

	userAccount := make([]*UserAccount, 0)

	qs := o.QueryTable(new(UserAccount))

	_, err := qs.All(&userAccount)
	return userAccount, err
}
