package spm

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/zp4rker/jpm/internal/spm/sock_api"
)

const SockAddr = "/tmp/spm/controller.sock"

type Controller struct {
	procs    []Proc
	procMap  map[net.Conn]Proc
	listener net.Listener
}

type Proc struct {
	Pid  int
	Conn net.Conn
	// TODO: Implement Proc struct
}

func (c *Controller) LaunchProc(cmdString string) error {
	cmdString = fmt.Sprintf("spm-wrapper '%v'", cmdString)
	cmd, err := CmdFromString(cmdString)
	if err != nil {
		return err
	}

	return cmd.Start()
}

func (c *Controller) SignalProc(pid int, signal os.Signal) error {
	osProc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return osProc.Signal(signal)
}

func (c *Controller) ReattachProc(pid int) error {
	return c.SignalProc(pid, SIGRQFB)
}

func (c *Controller) ProcList() *[]Proc {
	return &c.procs
}

func (c *Controller) Start() error {
	c.procMap = make(map[net.Conn]Proc)

	var err error
	var conn net.Conn
	for err == nil {
		conn, err = c.listener.Accept()
		err = c.AcceptConnection(conn)
	}

	return err
}

func (c *Controller) AcceptConnection(conn net.Conn) error {
	var proc Proc

	rd := bufio.NewReader(conn)
	var err error
	var input string
	for err == nil {
		input, err = rd.ReadString('\n')
		if err != nil {
			continue
		}

		req := sock_api.ParseInput(input)
		switch req {
		case sock_api.RegisterRequest:
			pid, err := strconv.Atoi(strings.Fields(input)[1])
			if err != nil {
				// should return error
				continue
			}
			proc = Proc{pid, conn}
			c.procs = append(c.procs, proc)
			c.procMap[conn] = proc
			fmt.Printf("Process with pid %v registered itself\n", pid)
		case sock_api.TerminatedInfo:
			fmt.Printf("Process with pid %v has terminated\n", proc.Pid)
			delete(c.procMap, conn)
		case sock_api.HeartbeatInfo:
			fmt.Printf("Received heartbeat from process with pid %v\n", proc.Pid)
		case sock_api.UnknownRequest:
			fmt.Println("Received unknown request!")
		}
	}

	return nil
}

func (c *Controller) InitSock() error {
	if err := os.MkdirAll("/tmp/spm", 0700); err != nil {
		return err
	}

	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		return err
	}

	c.listener = l
	return nil
}

func (c *Controller) CloseSock() error {
	err := c.listener.Close()
	if err != nil {
		return err
	}

	if err := os.RemoveAll(SockAddr); err != nil {
		return err
	}

	c.listener = nil
	return nil
}
