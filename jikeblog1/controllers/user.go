package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Profile() {

	c.Data["userid"] = "jike"
	c.Data["hobby"] = []string{"chess", "code"}

	c.TplNames = "user/profile.html"
}
