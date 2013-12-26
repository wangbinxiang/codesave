package main

import (
	"codesave/controllers"
	"codesave/libs"
	"github.com/astaxie/beego"
)

func init() {
	libs.MysqlRegisterDB()
}

func main() {

	beego.Router("/", &controllers.IndexController{})
	beego.Router("/a", &controllers.AskController{})
	beego.Router("/q", &controllers.QuestionController{})
	beego.Run()
}
