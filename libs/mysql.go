package libs

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var Orm orm.Ormer

func MysqlRegisterDB() {

	// mysqlUser := beego.AppConfig.String("mysqluser")
	// mysqlPass := beego.AppConfig.String("mysqlpass")
	// mysqlUrl := beego.AppConfig.String("mysqlurl")
	// mysqlDB := beego.AppConfig.String("mysqldb")
	// mysqlCharset := beego.AppConfig.String("mysqlcharset")

	orm.RegisterDataBase("default", "mysql", "root:@/codesave?charset=utf8")
	Orm = orm.NewOrm()
}

func MysqlRegisterModelWithPrefix(models ...interface{}) {
	for _, model := range models {
		orm.RegisterModelWithPrefix(beego.AppConfig.String("mysqlPrefix"), model)
	}

}
