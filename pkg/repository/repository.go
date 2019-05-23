package repository

import (
	"errors"

	"github.com/astaxie/beego/orm"
	log "github.com/sirupsen/logrus"
	"gitlab.com/bialas1993/socket-dispatcher/pkg/model"
)

var (
	ErrorCanNotFindRecord   = errors.New("repository: can not find record.")
	ErrorCanNotInsertRecord = errors.New("repository: can not insert record.")
	ErrorCanNotUpdateRecord = errors.New("repository: can not update record.")
)

func init() {
	orm.RegisterModel(new(model.Socket))
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	err := orm.RegisterDataBase("default", "sqlite3", "file:data.db")
	if err != nil {
		log.Panic(err)
	}
	orm.RunSyncdb("default", false, true)

}

type Repository interface {
	Insert(port int, hash string) bool
	Update(socket *model.Socket) bool
	FindSocket(hash string) (*model.Socket, error)
	FindPorts(start, end int) ([]*model.Socket, error)
}

type repo struct {
	db orm.Ormer
}

func New() Repository {
	orm.Debug = false

	return &repo{orm.NewOrm()}
}

func (r *repo) Insert(port int, hash string) bool {
	if _, err := r.db.Insert(&model.Socket{Port: port, Hash: hash}); err != nil {
		return false
	}

	return true
}

func (r *repo) FindSocket(hash string) (*model.Socket, error) {
	socket := model.Socket{Hash: hash}
	// todo: change to read or create
	if err := r.db.Read(&socket, "hash"); err != nil {
		return nil, ErrorCanNotFindRecord
	}

	log.Debugf("Finded socket on port %d by hash [%s]\n", socket.Port, hash)

	return &socket, nil
}

func (r *repo) FindPorts(start, end int) ([]*model.Socket, error) {
	// r.db.QueryTable("socket").Filter()

	// qs := r.db.QueryTable(&model.Socket{})

	// qs.Filter("socket__port__in", range(start, end)...)
	// condition := orm.NewCondition()

	// condition.And("socket__port__gte", start).And("socket__port__lte", end)
	//	cond := orm.NewCondition()
	//	cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)
	//	//sql-> WHERE T0.`profile_id` IS NOT NULL AND NOT T0.`Status` IN (?) OR T1.`age` >  2000
	//	num, err := qs.SetCond(cond1).Count()

	return nil, nil
}

func (r *repo) Update(socket *model.Socket) bool {
	if _, err := r.db.Update(socket); err == nil {
		return true
	}

	return false
}
