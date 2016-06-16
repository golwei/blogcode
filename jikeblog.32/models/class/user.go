package class

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     string `orm:"pk;size(30)"`
	Nick   string
	Info   string `orm:"null"`
	Hobby  string `orm:"null"`
	Email  string `orm:"unique"`
	Avator string `orm:"null"`
	Url    string `orm:"null"`

	Followers int
	Following int

	Regtime time.Time `orm:"auto_now_add"`

	Password string
	Private  int
}

const (
	PR_live = iota
	PR_login
	PR_post
)

const (
	//	0111
	DefaultPvt = 1<<3 - 1
)

func (u User) Get() *User {
	o := orm.NewOrm()
	err := o.Read(&u)
	if err == orm.ErrNoRows {
		return nil
	}
	return &u
}

//	CRUD
//	create
//	read
//	update
//	delete

//	read data form db
func (u *User) ReadDB() (err error) {
	o := orm.NewOrm()
	//	pk;overwrite
	if err = o.Read(u); err != nil {
		beego.Info(err)
	}
	return
}

//	create
func (u *User) Create() (err error) {
	//	new in pool
	o := orm.NewOrm()
	if _, err = o.Insert(u); err != nil {
		beego.Info(err)
	}
	return
}

//	update
func (u User) Update() (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(&u); err != nil {
		beego.Info(err)
	}
	return
}

//	delete
func (u User) Delete() (err error) {
	//	xxx1 & 1110 = xxx0
	//	~x ==> ^x ( -1 ^ x )
	u.Private &= ^1
	err = u.Update()
	return
}

func (u *User) ExistId() bool {
	o := orm.NewOrm()
	if err := o.Read(u, "Id"); err == orm.ErrNoRows {
		return false
	}
	return true
}

func (u *User) ExistEmail() bool {
	o := orm.NewOrm()
	return o.QueryTable("user").Filter("Email", u.Email).Exist()
}

func (u User) ExistEmailInUpdate() bool {
	o := orm.NewOrm()
	return o.QueryTable("user").Exclude("Id", u.Id).Filter("Email", u.Email).Exist()
}

func Testmodel() {

	u := User{
		Id:    "jike",
		Nick:  "jike",
		Email: "123@q.com",
	}

	o := orm.NewOrm()
	_, err := o.Insert(&u)

	u1 := User{Id: "jike"}
	err = o.Read(&u1)
	fmt.Println(u1)

	u2 := User{Nick: "jike"}
	err = o.Read(&u2, "nick")
	fmt.Println(u2)

	u2.Nick = "jike2"
	_, err = o.Update(&u2)
	u1 = User{Id: "jike"}
	err = o.Read(&u1)
	fmt.Println(u1)

	_, err = o.Delete(&u)
	u1 = User{Id: "jike"}
	err = o.Read(&u1)
	fmt.Println(u1)

	if err != nil {
		fmt.Println(err)
	}

}
