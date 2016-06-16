package controllers

import "jikeblog/models/class"

type UserController struct {
	//	beego.Controller
	BaseController
}

func (c *UserController) Profile() {

	id := c.Ctx.Input.Params[":id"]
	u := &class.User{Id: id}
	u.ReadDB()

	c.Data["u"] = u

	a := &class.Article{Author: u}
	as := a.Gets()

	c.Data["articles"] = as

	c.TplNames = "user/profile.html"
}

func (c *UserController) PageJoin() {
	c.TplNames = "user/join.html"
}

func (c *UserController) PageSetting() {
	c.CheckLogin()
	c.TplNames = "user/setting.html"
}
