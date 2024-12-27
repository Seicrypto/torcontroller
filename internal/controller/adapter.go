package controller

import (
	"bytes"
	"errors"
	"log"
	"net"
	"os/exec"

	"github.com/Seicrypto/torcontroller/internal/singleton/logger"
)

type CommandHandler struct {
	Logger        *log.Logger
	Socket        ConnectionAdapter // abstract socket interface
	CommandRunner CommandRunner     // CLI runner interface
}

// ConnectionAdapter for abstract socket behavior
type ConnectionAdapter interface {
	Dial() (net.Conn, error)
}

// RealSocket is the actual socket adapter.
type RealSocket struct {
	Address string
}

func (r *RealSocket) Dial() (net.Conn, error) {
	return net.Dial("tcp", r.Address)
}

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

func NewCommandHandler(socket ConnectionAdapter, runner CommandRunner, loggerAdapter *log.Logger) *CommandHandler {
	if loggerAdapter == nil {
		loggerAdapter = logger.GetLogger().Logger // Default to the global logger
	}
	return &CommandHandler{
		Logger:        loggerAdapter,
		Socket:        socket,
		CommandRunner: runner,
	}
}
