package repository

import (
	"github.com/astaxie/beego/orm"
	"gitlab.com/bialas1993/socket-dispatcher/pkg/model"
)

func init() {
	orm.RegisterModel(new(model.Socket), new(model.ProjectPorts))
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	err := orm.RegisterDataBase("default", "sqlite3", "file:data.db")
	if err != nil {
		panic(err)
	}
	orm.RunSyncdb("default", false, true)

}

type Repository interface {
}

type repo struct {
	db orm.Ormer
}

func New() Repository {
	orm.Debug = true

	return &repo{orm.NewOrm()}
}
