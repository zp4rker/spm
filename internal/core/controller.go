package core

import (
	"net"
	"os"
)

const SockAddr = "/tmp/spm/controller.sock"

func NewController() (*Controller, error) {
	return &Controller{}, nil
}

type Controller struct {
	procs []Proc
	listener net.Listener
}

type Proc struct {
	// TODO: Implement Proc struct
}

func (c *Controller) LaunchProc() {
	// TODO: Implement process launching
}

func (c *Controller) SignalProc(proc Proc, signal os.Signal) {
	// TODO: Implement process signalling
}

func (c *Controller) ReattachProc(pid int) {
	// TODO: Implement process reattaching
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