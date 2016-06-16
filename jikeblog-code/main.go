package main

import (
	"encoding/gob"
	_ "jikeblog/models"
	"jikeblog/models/class"
	_ "jikeblog/routers"
	"strings"

	"github.com/astaxie/beego"
)

func init() {

	gob.Register(class.User{})

	beego.AddFuncMap("split", SplitHobby)

}

func main() {
	beego.Run()
}

/*	Template Function	*/

func SplitHobby(s string, sep string) []string {
	return strings.Split(s, sep)
}
