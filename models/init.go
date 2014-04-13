package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var Orm orm.Ormer

func MysqlRegisterDB() {

	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPass := beego.AppConfig.String("mysqlpass")
	mysqlHost := beego.AppConfig.String("mysqlhost")
	mysqlPort := beego.AppConfig.String("mysqlport")
	mysqlDB := beego.AppConfig.String("mysqldb")

	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDB))
	orm.Debug = true
	Orm = orm.NewOrm()
}

func MysqlRegisterModelWithPrefix(models ...interface{}) {
	for _, model := range models {
		orm.RegisterModelWithPrefix(beego.AppConfig.String("mysqlPrefix"), model)
	}
}
