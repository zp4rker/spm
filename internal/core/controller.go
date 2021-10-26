package core

import (
	"fmt"
	"net"
	"os"
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

	return cmd.Run()
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

func (c *Controller) InitSock() error {
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