package class

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     string `orm:"pk"`
	Nick   string
	Info   string `orm:"null"`
	Hobby  string `orm:"null"`
	Email  string `orm:"unique"`
	Avator string `orm:"null"`
	Url    string `orm:"null"`

	Posts int

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
	DefaultPvt = 1<<3 - 1
)

//	CRUD
//	create
//	read
//	update
//	delete

func (u *User) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(u)
	return
}

func (u User) Create() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(&u)
	return
}

func (u User) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(&u)
	return
}

func (u User) Delete() (err error) {
	//	xxx1 & 1110 = xxx0
	//	~x ==> ^x (-1 ^ x)
	u.Private &= ^1
	err = u.Update()
	return
}

func (u User) ExistId() bool {
	o := orm.NewOrm()
	if err := o.Read(&u); err == orm.ErrNoRows {
		return false
	}
	return true
}

func (u User) ExistEmail() bool {
	o := orm.NewOrm()
	return o.QueryTable("user").Filter("Email", u.Email).Exist()
}
