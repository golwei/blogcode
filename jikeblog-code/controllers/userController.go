package controllers

import . "jikeblog/models/class"

type UserController struct {
	//	beego.Controller
	BaseController
}

func (c *UserController) Profile() {

	id := c.Ctx.Input.Params[":id"]
	u := &User{Id: id}
	u.ReadDB()

	c.Data["u"] = u

	as := Article{Author: u}.Gets()
	replys := Reply{Author: u}.Gets()

	c.Data["articles"] = as
	c.Data["replys"] = replys

	c.TplNames = "user/profile.html"
}

func (c *UserController) PageJoin() {
	c.TplNames = "user/join.html"
}

func (c *UserController) PageSetting() {
	c.CheckLogin()
	c.TplNames = "user/setting.html"
}
