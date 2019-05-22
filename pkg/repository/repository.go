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
	FindSocket(hash string) (*model.Socket, error)
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
	if err := r.db.Read(&socket, "hash"); err != nil {
		return nil, ErrorCanNotFindRecord
	}

	log.Debugf("Finded socket on port %d by hash [%s]\n", socket.Port, hash)

	return &socket, nil
}
