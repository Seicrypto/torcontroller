package controller

import (
	"bytes"
	"errors"
	"log"
	"net"
	"os/exec"

	"github.com/Seicrypto/torcontroller/internal/singleton/configuration"
	"github.com/Seicrypto/torcontroller/internal/singleton/logger"
)

type CommandHandler struct {
	Logger        *log.Logger
	Socket        ConnectionAdapter            // abstract socket interface
	CommandRunner CommandRunner                // CLI runner interface
	Config        *configuration.Configuration // Injected configuration
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

func NewCommandHandler(socket ConnectionAdapter, runner CommandRunner, log *log.Logger, cfg *configuration.Configuration) *CommandHandler {
	if log == nil {
		log = logger.GetLogger().Logger
	}
	if cfg == nil {
		cfg = configuration.GetConfig() // Default to singleton if not provided
	}
	return &CommandHandler{
		Logger:        log,
		Socket:        socket,
		CommandRunner: runner,
		Config:        cfg,
	}
}
