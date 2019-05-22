package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"gitlab.com/bialas1993/socket-dispatcher/pkg/socket"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Project configurated")
	log.Info("Hello world")

	s := socket.New(3000)

	fmt.Printf("%+v\n", s.IsLocked())
	s.Open()
	fmt.Printf("%+v\n", s.IsLocked())
	// s.Close()
	fmt.Printf("%+v\n", s.IsLocked())

	// repository.New()
}
