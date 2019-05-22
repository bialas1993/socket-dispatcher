package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"gitlab.com/bialas1993/socket-dispatcher/pkg/process"
	"gitlab.com/bialas1993/socket-dispatcher/pkg/repository"
)

const (
	socketDispatcherPorts = "SOCKET_DISPATCHER_PORTS"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	port := 8000
	branch := "asd"

	h := md5.New()
	io.WriteString(h, branch)

	hash := fmt.Sprintf("%x", h.Sum(nil))
	log.Debugf("Branch hash: %s", hash)

	portRange := os.Getenv(socketDispatcherPorts)
	if len(portRange) > 0 {
		ports := strings.Split(portRange, "-")
		portStart, err := strconv.Atoi(ports[0])
		if err != nil {
			log.Errorf("cmd: can not parse port range")
			os.Exit(1)
			return
		}

		portEnd, err := strconv.Atoi(ports[1])
		if err != nil {
			log.Errorf("cmd: can not parse port range")
			os.Exit(1)
			return
		}

		log.Printf("Ports: %d-%d", portStart, portEnd)

		pid, _ := process.FindPidByPort(port)
		log.Println(process.Kill(pid))
		repo := repository.New()
		repo.FindSocket(hash)
	}

	log.Error("cmd: Port range is not set!")
	os.Exit(1)
}
