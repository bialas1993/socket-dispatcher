package model

import "time"

type Socket struct {
	Id      int       `orm:"auto"`
	Hash    string    `orm:"size(16)"`
	Port    int       `orm:"unique"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
