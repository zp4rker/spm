package spm

import (
	"os/exec"

	"github.com/google/shlex"
)

func CmdFromString(cmdString string) (*exec.Cmd, error) {
	cmdSplit, err := shlex.Split(cmdString)
	if err != nil {
		return nil, err
	}

	return exec.Command(cmdSplit[0], cmdSplit[1:]...), nil
}
