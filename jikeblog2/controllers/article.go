package controllers

import (
	"jikeblog/models/class"
	"strconv"
)
import . "fmt"

type ArticleController struct {
	BaseController
	ret RET
}

func (c *ArticleController) PageNew() {
	c.CheckLogin()
	c.TplNames = "article/new.html"
}

func (c *ArticleController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Params[":id"])
	a := &class.Article{Id: id}
	a.ReadDB()
	a.Author.ReadDB()
	c.Data["article"] = a
	c.TplNames = "article/article.html"
}

func (c *ArticleController) PageEdit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Params[":id"])
	a := &class.Article{Id: id}
	a.ReadDB()
	a.Author.ReadDB()
	c.Data["article"] = a
	c.TplNames = "article/edit.html"
}

func (c *ArticleController) Edit() {
	c.CheckLogin()
	u := c.GetSession("user").(class.User)

	id, _ := strconv.Atoi(c.Ctx.Input.Params[":id"])
	a := &class.Article{Id: id}
	a.ReadDB()

	if u.Id != a.Author.Id {
		c.DoLogout()
	}

	a.Title = c.GetString("title")
	a.Content = c.GetString("content")

	a.Update()

	c.ret.Ok = true
	c.Data["json"] = c.ret
	c.ServeJson()

}

func (c *ArticleController) New() {
	c.CheckLogin()

	u := c.GetSession("user").(class.User)

	a := &class.Article{
		Title:   c.GetString("title"),
		Content: c.GetString("content"),
		Author:  &u,
	}

	n, err := a.Create()

	if err == nil {
		c.ret.Ok = true
		c.ret.Content = n
		c.Data["json"] = c.ret
		c.ServeJson()
		return
	}

	c.ret.Content = Sprint(err)

	c.Data["json"] = c.ret
	c.ServeJson()
}

func (c *ArticleController) Del() {
	c.CheckLogin()
	u := c.GetSession("user").(class.User)

	id, _ := strconv.Atoi(c.Ctx.Input.Params[":id"])
	a := &class.Article{Id: id}
	a.ReadDB()

	if u.Id != a.Author.Id {
		c.DoLogout()
	}

	a.Defunct = true
	a.Update()

	c.Redirect("/user/"+a.Author.Id, 302)
}
