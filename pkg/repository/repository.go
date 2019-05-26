package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/astaxie/beego/orm"
	log "github.com/sirupsen/logrus"
	"gitlab.com/bialas1993/socket-dispatcher/pkg/model"
)

var (
	ErrorCanNotFindRecord   = errors.New("repository: can not find record.")
	ErrorCanNotFindRecords  = errors.New("repository: can not find records.")
	ErrorCanNotInsertRecord = errors.New("repository: can not insert record.")
	ErrorCanNotUpdateRecord = errors.New("repository: can not update record.")
)

func init() {
	orm.RegisterModel(new(model.Socket))
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

}

type Repository interface {
	Insert(port int, hash string) bool
	Update(socket *model.Socket) bool
	FindSocketHash(hash string) (*model.Socket, error)
	FindSocketPorts(start, end int) ([]*model.Socket, error)
}

type repo struct {
	db orm.Ormer
}

type config interface {
	DatabasePath() string
}

func New(cfg config) Repository {
	err := orm.RegisterDataBase("default", "sqlite3",
		fmt.Sprintf("file:%s/data.db", strings.TrimRight(cfg.DatabasePath(), "/")))
	if err != nil {
		log.Panic(err)
	}
	orm.RunSyncdb("default", false, false)

	return &repo{orm.NewOrm()}
}

func (r *repo) Insert(port int, hash string) bool {
	if _, err := r.db.Insert(&model.Socket{Port: port, Hash: hash}); err != nil {
		return false
	}

	return true
}

func (r *repo) FindSocketHash(hash string) (*model.Socket, error) {
	socket := model.Socket{Hash: hash}
	// todo: change to read or create
	if err := r.db.Read(&socket, "hash"); err != nil {
		return nil, ErrorCanNotFindRecord
	}

	log.Debugf("Finded socket on port %d by hash [%s]\n", socket.Port, hash)

	return &socket, nil
}

func (r *repo) FindSocketPorts(start, end int) ([]*model.Socket, error) {
	var data []*model.Socket
	var ports []interface{}

	qs := r.db.QueryTable(&model.Socket{})

	for i := start; i <= end; i++ {
		ports = append(ports, i)
	}

	if _, err := qs.Filter("port__in", ports...).OrderBy("updated").All(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func (r *repo) Update(socket *model.Socket) bool {
	if _, err := r.db.Update(socket); err == nil {
		return true
	}

	return false
}
