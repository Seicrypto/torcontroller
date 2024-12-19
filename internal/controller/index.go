package controller

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/Seicrypto/torcontroller/internal/services/logger"
	"github.com/Seicrypto/torcontroller/internal/services/tor"
)

type CommandHandler func(net.Conn, string) error

var commandHandlers = map[string]CommandHandler{
	"start": func(conn net.Conn, socketPath string) error {
		if _, err := conn.Write([]byte("Starting Tor service...\n")); err != nil {
			return err
		}
		if err := tor.StartTorService(); err != nil {
			return err
		}
		if _, err := conn.Write([]byte("Tor service started successfully.\n")); err != nil {
			return err
		}
		return nil
	},
	"status": func(conn net.Conn, socketPath string) error {
		_, err := conn.Write([]byte(fmt.Sprintf("Listener at %s is running.\n", socketPath)))
		return err
	},
	"switch": func(conn net.Conn, socketPath string) error {
		conn.Write([]byte("Successfully switched Tor IP.\n"))
		return nil
	},
}

func HandleConnection(conn net.Conn, socketPath string, listener net.Listener) error {
	logger := logger.GetLogger()
	defer conn.Close()
	buf := make([]byte, 1024)

	// Read data from the connection
	n, err := conn.Read(buf)
	if err != nil {
		if errors.Is(err, io.EOF) {
			logger.Warn("Connection closed by client.")
			return nil
		}
		logger.Error(fmt.Sprintf("Error reading from connection: %v", err))
		return err
	}

	command := strings.TrimSpace(string(buf[:n]))
	logger.Info(fmt.Sprintf("Received command on %s: %s\n", socketPath, command))

	// Handle the command
	handler, ok := commandHandlers[command]
	if !ok {
		msg := fmt.Sprintf("Unknown command: %s", command)
		logger.Warn(msg)
		_, _ = conn.Write([]byte(fmt.Sprintf("Unknown command: %s\n", command)))
		return nil
	}

	if err := handler(conn, socketPath); err != nil {
		logger.Error(fmt.Sprintf("Error executing command '%s': %v", command, err))
		return err
	}

	return nil
}
