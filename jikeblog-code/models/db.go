package models

import (
	"fmt"
	"log"

	. "jikeblog/models/class"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	orm.Debug, _ = beego.AppConfig.Bool("DB::debug")

	switch beego.AppConfig.String("DB::db") {

	case "mysql":
		dbu := beego.AppConfig.String("DB::user")
		dbp := beego.AppConfig.String("DB::pass")
		dbn := beego.AppConfig.String("DB::name")

		orm.RegisterDriver("mysql", orm.DR_MySQL)
		orm.RegisterDataBase(
			"default",
			"mysql",
			fmt.Sprintf(`%s:%s@/%s?charset=utf8&loc=%s`, dbu, dbp, dbn, `Asia%2FShanghai`),
			10,
			10,
		)

	case "sqlite":
		orm.RegisterDriver("sqlite3", orm.DR_Sqlite)
		orm.RegisterDataBase("default", "sqlite3", beego.AppConfig.String("DB::sqlitefile"))

	default:
		log.Fatalln(beego.AppConfig.String("DB::db"))
	}

	orm.RegisterModel(new(User), new(Article), new(Tag), new(Reply))

	_ = orm.RunSyncdb("default", false, true)

}
