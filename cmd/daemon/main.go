package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/caarlos0/env"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"gitlab.com/bialas1993/socket-dispatcher/pkg/repository"
)

// Config for app
type config struct {
	SocketDispatcherPorts string `env:"SOCKET_DISPATCHER_PORTS,required"`
}

var (
	branch string
	cfg    config
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
		return
	}

	flag.StringVar(&branch, "branch", "", "branch name")
	flag.Parse()

	if len(branch) == 0 {
		log.Error("Branch name is not defined.")
		os.Exit(1)
		return
	}
}

func main() {
	h := md5.New()
	io.WriteString(h, branch)

	hash := fmt.Sprintf("%x", h.Sum(nil))
	log.Debugf("Branch hash: %s", hash)

	if len(cfg.SocketDispatcherPorts) > 0 {
		ports := strings.Split(cfg.SocketDispatcherPorts, "-")
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

		log.Debugf("Ports: %d-%d", portStart, portEnd)

		// pid, _ := process.FindPidByPort(port)
		// log.Println(process.Kill(pid))
		repo := repository.New()
		socket, err := repo.FindSocket(hash)
		if err == nil {
			// todo: // update time, because is usage

			log.Printf("socket updated: %#v", socket)
			if ok := repo.Update(socket); ok {
				println()
				print("PORT:")
				println(socket.Port)

				// kill pid ?
				return
			}
			log.Error("Can not update socket row.")
			os.Exit(2)
			return
		}

		repo.FindPorts(portStart, portEnd)

		// todo: insert new

		// repo.Insert(8001, hash)
		return
	}

	log.Error("cmd: Port range is not set!")
	os.Exit(1)
}
