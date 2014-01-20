package main

import (
	"codesave/models"
	_ "codesave/routers"
	"github.com/astaxie/beego"
)

func init() {
	models.MysqlRegisterDB()
}

func main() {
	beego.Run()
}
