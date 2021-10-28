package spm

import (
	"bufio"
	"fmt"
	"github.com/zp4rker/jpm/internal/spm/sock_api"
	"net"
	"os"
	"strconv"
	"strings"
)

const SockAddr = "/tmp/spm/controller.sock"

type Controller struct {
	procs []Proc
	listener net.Listener
}

type Proc struct {
	Pid int
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

func (c *Controller) Start() {
	var err error
	var conn net.Conn
	for err == nil {
		conn, err = c.listener.Accept()
		err = c.AcceptConnection(conn)
	}
}

func (c *Controller) AcceptConnection(conn net.Conn) error {
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
			c.procs = append(c.procs, Proc{pid, conn})
			fmt.Printf("Process with pid %v registered itself", pid)
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