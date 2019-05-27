package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"gitlab.com/bialas1993/socket-dispatcher/pkg/model"
	"gitlab.com/bialas1993/socket-dispatcher/pkg/repository"
)

const (
	ExitNotSetVariabled = iota + 127
	ExitCanNotRead
	ExitCanNotUpdate
)

var (
	branch string
	debug  bool
	kill   bool
	cfg    config
)

// Config for app
type config struct {
	SocketDispatcherPorts string `env:"SOCKET_DISPATCHER_PORTS,required"`
	DbPath                string `env:"DATABASE_PATH" envDefault:"."`
}

func (c config) DatabasePath() string {
	return c.DbPath
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.ErrorLevel)

	godotenv.Load()
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
		return
	}

	flag.StringVar(&branch, "branch", "", "branch name")
	flag.BoolVar(&kill, "kill", false, "kill process which one use selected port")
	flag.BoolVar(&debug, "debug", false, "enable debug")

	flag.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
	}

	if len(branch) == 0 {
		log.Error("Branch name is not defined.")
		os.Exit(ExitNotSetVariabled)
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
			log.Errorf("cmd: can not parse port range '%s'", cfg.SocketDispatcherPorts)
			os.Exit(ExitNotSetVariabled)
			return
		}

		portEnd, err := strconv.Atoi(ports[1])
		if err != nil {
			log.Errorf("cmd: can not parse port range")
			os.Exit(ExitNotSetVariabled)
			return
		}

		log.Debugf("Ports: %d-%d", portStart, portEnd)

		repo := repository.New(cfg)
		socket, err := repo.FindSocketHash(hash)
		if err == nil {
			log.Debugf("socket updated: %#v", socket) // is used
			if ok := repo.Update(socket); ok {
				printPort(socket.Port)
				return
			}

			log.Error("Can not update socket row.")
			os.Exit(ExitCanNotUpdate)
			return
		}

		sockets, err := repo.FindSocketPorts(portStart, portEnd)
		if err != nil {
			log.WithError(err).Error("main: can not find socket for port range.")
			os.Exit(ExitCanNotRead)
		}

		// find free port
		var port int
		if ((portEnd + 1) - portStart) > len(sockets) {
			for i := portStart; i <= portEnd; i++ {
				if isFreePort(i, sockets) {
					port = i
					log.Debugf("Free port: %d", port)
					break
				}
			}

			repo.Insert(port, hash)
			printPort(port)
			return
		}

		if len(sockets) > 0 {
			socket := sockets[0]
			socket.Hash = hash

			if ok := repo.Update(socket); !ok {
				log.Error("main: can not update socket use info")
				os.Exit(ExitCanNotUpdate)
			}

			printPort(socket.Port)
		}

		return
	}

	log.Error("cmd: Port range is not set!")
	os.Exit(ExitNotSetVariabled)
}

func printPort(port int) {
	log.Debugf("Port: %d", port)
	println(port)
}

func isFreePort(port int, sockets []*model.Socket) bool {
	for _, socket := range sockets {
		if socket.Port == port {
			return false
		}
	}

	return true
}
