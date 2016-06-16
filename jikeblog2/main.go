package main

import (
	"encoding/gob"
	_ "jikeblog/models"
	"jikeblog/models/class"
	_ "jikeblog/routers"

	"github.com/astaxie/beego"
)

func init() {
	gob.Register(class.User{})
}

func main() {
	beego.Run()
}
