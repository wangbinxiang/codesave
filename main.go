package main

import (
	"github.com/astaxie/beego"
	"github.com/wangbinxiang/codesave/models"
	_ "github.com/wangbinxiang/codesave/routers"
)

func init() {
	models.MysqlRegisterDB()
}

func main() {
	beego.Run()
}
