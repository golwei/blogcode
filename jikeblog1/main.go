package main

import (
	_ "jikeblog/models"
	_ "jikeblog/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
