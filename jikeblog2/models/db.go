package models

import (
	"fmt"
	"jikeblog/models/class"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.Debug = true

	switch beego.AppConfig.String("DB::db") {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DR_MySQL)
		orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&loc=%s",
			beego.AppConfig.String("DB::user"),
			beego.AppConfig.String("DB::pass"),
			beego.AppConfig.String("DB::name"),
			`Asia%2FShanghai`,
		))
	case "sqlite":
		orm.RegisterDriver("sqlite", orm.DR_Sqlite)
		orm.RegisterDataBase("default", "sqlite3", beego.AppConfig.String("DB::file"))
	}

	orm.RegisterModel(new(class.User), new(class.Article))

	orm.RunSyncdb("default", false, true)
}
