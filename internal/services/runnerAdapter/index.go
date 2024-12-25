package runneradapter

import (
	"bytes"
	"errors"
	"os/exec"
)

type CommandRunner interface {
	Run(name string, args ...string) (string, error)
}

type RealCommandRunner struct{}

func (r *RealCommandRunner) Run(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf

	err := cmd.Run()
	if err != nil {
		return "", errors.New(errBuf.String())
	}
	return out.String(), nil
}
