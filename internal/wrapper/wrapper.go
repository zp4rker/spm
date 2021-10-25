package wrapper

import (
	"errors"
	"github.com/google/shlex"
	"io"
	"os/exec"
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
	cmd *exec.Cmd
	stdin io.WriteCloser
	stdout, stderr io.ReadCloser
}

func (wrapper *Wrapper) Run() error {
	return wrapper.cmd.Run()
}

func (wrapper *Wrapper) WriteStdin(input string) error {
	_, err := wrapper.stdin.Write([]byte(input))
	return err
}

func (wrapper *Wrapper) ReadStdout() (string, error) {
	bytes, err := io.ReadAll(wrapper.stdout)
	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}

func (wrapper *Wrapper) ReadSterr() (string, error) {
	bytes, err := io.ReadAll(wrapper.stderr)
	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}