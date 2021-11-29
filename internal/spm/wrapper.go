package spm

import (
	"errors"
	"io"
	"net"
	"os/exec"
	"time"

	"github.com/google/shlex"
)

func NewWrapper(cmdString string) (*Wrapper, error) {
	cmdSplit, err := shlex.Split(cmdString)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(cmdSplit[0], cmdSplit[1:]...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, errors.New("unable to create stdin pipe")
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.New("unable to create stdout pipe")
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, errors.New("unable to create stderr pipe")
	}

	return &Wrapper{
		cmd:    cmd,
		stdin:  stdin,
		stdout: stdout,
		stderr: stderr,
	}, nil
}

type Wrapper struct {
	cmd            *exec.Cmd
	stdin          io.WriteCloser
	stdout, stderr io.ReadCloser
}

func (w *Wrapper) Run() error {
	return w.cmd.Run()
}

func (w *Wrapper) Start() error {
	return w.cmd.Start()
}

func (w *Wrapper) Wait() error {
	return w.cmd.Wait()
}

func StartHeartbeat(conn net.Conn) {
	for {
		_, _ = conn.Write([]byte("HEARTBEAT\n"))
		time.Sleep(5 * time.Second)
	}
}

func (w *Wrapper) WriteStdin(input string) error {
	_, err := w.stdin.Write([]byte(input))
	return err
}

func (w *Wrapper) ReadStdout() (string, error) {
	bytes, err := io.ReadAll(w.stdout)
	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}

func (w *Wrapper) ReadSterr() (string, error) {
	bytes, err := io.ReadAll(w.stderr)
	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}
