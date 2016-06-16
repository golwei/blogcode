package routers

import (
	"jikeblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/join", &controllers.UserController{}, `get:PageJoin`)
	beego.Router("/login", &controllers.UserController{}, `post:Login`)
	beego.Router("/register", &controllers.UserController{}, `post:Register`)
	beego.Router("/logout", &controllers.UserController{}, `get:Logout`)

	beego.Router("/setting", &controllers.UserController{}, `get:PageSetting;post:Setting`)

	beego.Router("/article/new", &controllers.ArticleController{}, `get:PageNew;post:New`)
	beego.Router("/article/edit/:id([0-9]+)", &controllers.ArticleController{}, `get:PageEdit;post:Edit`)
	beego.Router("/article/del/:id([0-9]+)", &controllers.ArticleController{}, `get:Del`)

	beego.Router("/article/:id([0-9]+)", &controllers.ArticleController{}, `get:Get`)

	beego.Router("/archive", &controllers.ArticleController{}, `get:Archive`)

	beego.Router("/reply/new", &controllers.ReplyController{}, `post:New`)

	beego.Router("/user/:id", &controllers.UserController{}, `get:Profile`)
	beego.Router("/user/setting", &controllers.UserController{}, `get:Profile`)
	beego.Router("/api/user/profile", &controllers.UserController{}, `get:API_Profile`)

}
