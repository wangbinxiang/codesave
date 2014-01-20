package routers

import (
	"codesave/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/:page:int", &controllers.IndexController{})

	beego.Router("/a", &controllers.AskController{})
	beego.Router("/a/:qid:int", &controllers.AskController{})

	beego.Router("/q/:qid:int", &controllers.QuestionController{})
	beego.Router("/q/c", &controllers.QuestionController{}, "*:GetComment")

	beego.Router("/r", &controllers.RegisterController{})
	beego.Router("/r/verify", &controllers.RegisterController{}, "*:Verify")

	beego.Router("/l", &controllers.LoginController{})

	beego.Router("/o", &controllers.LogoutController{})

	beego.Router("/c", &controllers.CommentController{})

	beego.Router("/u", &controllers.UserController{})
}
