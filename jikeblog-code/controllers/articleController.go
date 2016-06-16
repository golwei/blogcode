package controllers

import (
	"fmt"
	. "jikeblog/models/class"
	"strconv"
	"strings"
)

type ArticleController struct {
	BaseController
	ret RET
}

func (c *ArticleController) Archive() {

	errmsg := ""

	a := Article{}
	if len(c.GetString("tag")) > 0 {
		tag := Tag{Name: c.GetString("tag")}.Get()
		if tag == nil {
			errmsg += fmt.Sprintf("Tag[%s] is not exist.\n", c.GetString("tag"))
		} else {
			a.Tags = []*Tag{tag}
		}
	}

	if len(c.GetString("author")) > 0 {
		author := User{Id: c.GetString("author")}.Get()
		if author == nil {
			errmsg += fmt.Sprintf("User[%s] is not exist.\n", c.GetString("author"))
		} else {
			a.Author = author
		}
	}

	if len(errmsg) == 0 {
		rets := a.Gets()
		c.Data["articles"] = rets
	}

	c.Data["err"] = errmsg

	c.TplNames = "article/archive.html"

}

func (c *ArticleController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Params[":id"])
	a := &Article{Id: id}
	a.ReadDB()
	a.Author.ReadDB()

	a.Replys = Reply{Article: a}.Gets()
	c.Data["article"] = a
	c.Data["replyTree"] = a.GetReplyTree()

	c.TplNames = "article/article.html"
}

func (c *ArticleController) PageEdit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Params[":id"])
	a := &Article{Id: id}
	a.ReadDB()
	a.Author.ReadDB()
	c.Data["article"] = a
	c.TplNames = "article/edit.html"
}

func (c *ArticleController) Edit() {
	c.CheckLogin()
	u := c.GetSession("user").(User)

	id, _ := strconv.Atoi(c.Ctx.Input.Params[":id"])
	a := &Article{Id: id}
	a.ReadDB()

	if u.Id != a.Author.Id {
		c.DoLogout()
	}

	strs := strings.Split(c.GetString("tag"), ",")
	tags := []*Tag{}
	for _, v := range strs {
		tags = append(tags, Tag{Name: strings.TrimSpace(v)}.GetOrNew())
	}
	a.Title = c.GetString("title")
	a.Content = c.GetString("content")
	a.Tags = tags

	a.Update()

	c.ret.Ok = true
	c.Data["json"] = c.ret
	c.ServeJson()

}

func (c *ArticleController) PageNew() {
	c.CheckLogin()
	c.TplNames = "article/new.html"
}
func (c *ArticleController) New() {
	c.CheckLogin()

	u := c.GetSession("user").(User)

	a := &Article{
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

	c.ret.Content = err.Error()

	c.Data["json"] = c.ret
	c.ServeJson()
}

func (c *ArticleController) Del() {
	c.CheckLogin()
	u := c.GetSession("user").(User)

	id, _ := strconv.Atoi(c.Ctx.Input.Params[":id"])
	a := &Article{Id: id}
	a.ReadDB()

	if u.Id != a.Author.Id {
		c.DoLogout()
	}

	a.Defunct = true
	a.Update()

	c.Redirect("/user/"+a.Author.Id, 302)
}
