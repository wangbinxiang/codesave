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
	beego.Router("/a/:qid:int", &controllers.AskController{})

	beego.Router("/q/:qid:int", &controllers.QuestionController{})

	beego.Router("/r", &controllers.RegisterController{})
	beego.Router("/r/verify", &controllers.RegisterController{}, "*:Verify")
	beego.Run()
}
