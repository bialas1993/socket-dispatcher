package socket

import (
	"fmt"
	"net"
	"time"
)

const (
	connectionTypeTCP = "tcp"
	connectionTypeUDP = "udp"
	timeTimeout       = 5
	hostPattern       = "127.0.0.1:%d"
)

type Socket interface {
	IsLocked() bool
	Open() bool
	Close()
}

type socket struct {
	port     uint
	listener net.Conn
}

// New function for validate socket is available.
func New(port uint) Socket {
	return &socket{port: port}
}

func (s *socket) IsLocked() bool {
	l, err := net.DialTimeout(connectionTypeTCP, fmt.Sprintf(hostPattern, s.port), time.Duration(timeTimeout)*time.Second)
	if err != nil {
		return true
	}
	l.Close()

	return false
}

func (s *socket) Open() bool {
	var err error
	if s.listener, err = net.DialTimeout(connectionTypeTCP, fmt.Sprintf(hostPattern, s.port), time.Duration(timeTimeout)*time.Second); err != nil {
		return false
	}

	return true
}

func (s *socket) Close() {
	s.listener.Close()
}
