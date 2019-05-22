package process

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/cakturk/go-netstat/netstat"
)

var ErrorCanNotFindPid = errors.New("process: can not find process by port.")

func FindPidByPort(port int) (int, error) {
	socks, _ := netstat.TCPSocks(func(s *netstat.SockTabEntry) bool {
		return s.State == netstat.Listen && s.LocalAddr.Port == uint16(port)
	})

	if len(socks) > 0 {
		for _, sock := range socks {
			process := strings.Split(sock.Process.String(), "/")
			pid, _ := strconv.Atoi(process[0])
			return pid, nil
		}
	}

	return 0, ErrorCanNotFindPid
}

func Kill(pid int) error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return ErrorCanNotFindPid
	}

	return p.Signal(os.Kill)
}
