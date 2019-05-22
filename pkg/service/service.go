package service

import (
	"errors"

	"github.com/takama/daemon"
)

const (
	daemonPort = 6000
)

var (
	ErrorCanNotCreate = errors.New("service: Can not create daemon instance.")
)

type Daemon interface {
	Do() error
}

type daemonService struct {
	daemon.Daemon
}

func NewDaemon(name, description string) (Daemon, error) {
	srv, err := daemon.New(name, description)
	if err != nil {
		return nil, ErrorCanNotCreate
	}

	return &daemonService{srv}, nil
}

func (s *daemonService) Do() error {
	return nil
}
