package controller_test

import (
	"bytes"
	"log"
	"net"

	"github.com/Seicrypto/torcontroller/internal/controller"
)

// MockSocket is a mock implementation of ConnectionAdapter.
type MockSocket struct {
	Response []byte
	Error    error
}

func (m *MockSocket) Dial() (net.Conn, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	client, server := net.Pipe() // Simulate a connection
	go func() {
		server.Write(m.Response)
		server.Close()
	}()
	return client, nil
}

// MockCommandRunner is a mock implementation of CommandRunner.
type MockCommandRunner struct {
	Output string
	Error  error
}

func (m *MockCommandRunner) Run(name string, args ...string) (string, error) {
	if m.Error != nil {
		return "", m.Error
	}
	return m.Output, nil
}

// NewMockCommandHandler creates a new CommandHandler with mocked dependencies.
func NewMockCommandHandler(loggerAdapter *log.Logger, socketResponse []byte, socketError error, runnerOutput string, runnerError error) *controller.CommandHandler {
	if loggerAdapter == nil {
		loggerAdapter = log.New(&bytes.Buffer{}, "TEST: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	return &controller.CommandHandler{
		Logger: loggerAdapter,
		Socket: &MockSocket{
			Response: socketResponse,
			Error:    socketError,
		},
		CommandRunner: &MockCommandRunner{
			Output: runnerOutput,
			Error:  runnerError,
		},
	}
}
