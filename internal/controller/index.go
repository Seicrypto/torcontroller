package controller

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

// type CommandHandler func(net.Conn, string) error

// var commandHandlers = map[string]CommandHandler{
// "start": func(conn net.Conn, socketPath string) error {
// 	logger := logger.GetLogger()
// 	if err := tor.StartTorService(); err != nil {
// 		logger.Error(fmt.Sprintf("Failed to start Tor service: %v", err))
// 		_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
// 		return err
// 	}
// 	if err := privoxy.StartPrivoxyService(); err != nil {
// 		logger.Error(fmt.Sprintf("Failed to start Privoxy service: %v", err))
// 		_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
// 		return err
// 	}
// 	if err := iptable.ApplyIptablesRules(); err != nil {
// 		logger.Error(fmt.Sprintf("Failed to apply iptables rules: %v", err))
// 		_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
// 		return err
// 	}
// 	if _, err := conn.Write([]byte("done\n")); err != nil {
// 		logger.Error(fmt.Sprintf("Failed to send final response: %v", err))
// 		return err
// 	}
// 	logger.Info("Tor service started successfully.\n")
// 	return nil
// },
// 	"traffic": func(conn net.Conn, socketPath string) error {
// 		logger := logger.GetLogger()
// 		readTraffic, writtenTraffic, err := tor.GetTorTrafficMetrics()
// 		if err != nil {
// 			logger.Error(fmt.Sprintf("Error fetching traffic metrics: %v\n", err))
// 			fmt.Fprintf(conn, "Error: %v\n", err)
// 			return err
// 		}
// 		response := fmt.Sprintf("Traffic Read: %s bytes, Traffic Written: %s bytes\n", readTraffic, writtenTraffic)
// 		if _, err := conn.Write([]byte(response)); err != nil {
// 			logger.Error(fmt.Sprintf("Failed to send traffic response: %v", err))
// 			return err
// 		}
// 		logger.Info("Get traffic successfully")
// 		return nil
// 	},
// 	"status": func(conn net.Conn, socketPath string) error {
// 		_, err := conn.Write([]byte(fmt.Sprintf("Listener at %s is running.\n", socketPath)))
// 		return err
// 	},
// 	"switch": func(conn net.Conn, socketPath string) error {
// 		logger := logger.GetLogger()
// 		if err := tor.SwitchTorCircuit(); err != nil {
// 			logger.Error(fmt.Sprintf("Failed to switch Tor Circuit: %v", err))
// 			_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
// 			return err
// 		}
// 		if _, err := conn.Write([]byte("done\n")); err != nil {
// 			logger.Error(fmt.Sprintf("Failed to send final response: %v", err))
// 			return err
// 		}
// 		return nil
// 	},
// 	"stop": func(conn net.Conn, socketPath string) error {
// 		logger := logger.GetLogger()
// 		if err := iptable.ClearIptablesRules(); err != nil {
// 			logger.Error(fmt.Sprintf("Failed to clear iptables rules: %v", err))
// 			_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
// 			return err
// 		}
// 		if err := privoxy.StopPrivoxyService(); err != nil {
// 			logger.Error(fmt.Sprintf("Failed to stop Privoxy service: %v", err))
// 			_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
// 			return err
// 		}
// 		if err := tor.StopTorService(); err != nil {
// 			logger.Error(fmt.Sprintf("Failed to stop Tor service: %v", err))
// 			_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
// 			return err
// 		}
// 		if _, err := conn.Write([]byte("done\n")); err != nil {
// 			logger.Error(fmt.Sprintf("Failed to send final response: %v", err))
// 			return err
// 		}
// 		logger.Info("Tor service stopped successfully")
// 		return nil
// 	},
// }

// func HandleConnection(conn net.Conn, socketPath string, listener net.Listener) error {
// 	logger := logger.GetLogger()
// 	defer conn.Close()
// 	buf := make([]byte, 1024)

// 	// Read data from the connection
// 	n, err := conn.Read(buf)
// 	if err != nil {
// 		if errors.Is(err, io.EOF) {
// 			logger.Warn("Connection closed by client.")
// 			return nil
// 		}
// 		logger.Error(fmt.Sprintf("Error reading from connection: %v", err))
// 		return err
// 	}

// 	command := strings.TrimSpace(string(buf[:n]))
// 	logger.Info(fmt.Sprintf("Received command on %s: %s\n", socketPath, command))

// 	// Handle the command
// 	handler, ok := commandHandlers[command]
// 	if !ok {
// 		msg := fmt.Sprintf("Unknown command: %s", command)
// 		logger.Warn(msg)
// 		_, _ = conn.Write([]byte(fmt.Sprintf("Unknown command: %s\n", command)))
// 		return nil
// 	}

// 	if err := handler(conn, socketPath); err != nil {
// 		logger.Error(fmt.Sprintf("Error executing command '%s': %v", command, err))
// 		return err
// 	}

// 	return nil
// }

func HandleConnection(conn net.Conn, socketPath string, listener net.Listener) error {
	handler := NewCommandHandler(&RealSocket{Address: "127.0.0.1:9051"}, &RealCommandRunner{}, nil)
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		if errors.Is(err, io.EOF) {
			handler.Logger.Println("[WARN] Connection closed by client.")
			return nil
		}
		handler.Logger.Printf("[ERROR] Error reading from connection: %v", err)
		return err
	}

	command := strings.TrimSpace(string(buf[:n]))
	handler.Logger.Printf("[INFO] Received command on %s: %s", socketPath, command)

	switch command {
	case "start":
		if err := handler.StartTorService(); err != nil {
			handler.Logger.Printf("[ERROR] %v", err)
			_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
			return err
		}
		if err := handler.StartPrivoxyService(); err != nil {
			handler.Logger.Printf("[ERROR] %v", err)
			_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
			return err
		}
		if err := handler.ApplyIptablesIPv4(); err != nil {
			handler.Logger.Printf("[ERROR] %v", err)
			_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
			return err
		}
		_, _ = conn.Write([]byte("Done\n"))
		handler.Logger.Println("[INFO] Tor service started successfully.")
		return nil
	case "switch":
		if err := handler.SwitchTorCircuit(); err != nil {
			handler.Logger.Printf("[ERROR] Failed to switch Tor circuit: %v", err)
			_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
			return err
		}
		_, _ = conn.Write([]byte("Circuit switched successfully\n"))
		handler.Logger.Println("[INFO] Successfully switched Tor circuit.")
		return nil
	case "traffic":
		readTraffic, writtenTraffic, err := handler.GetTorTrafficMetrics()
		if err != nil {
			handler.Logger.Printf("[ERROR] %v", err)
			_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
			return err
		}
		response := fmt.Sprintf("Traffic Read: %s bytes, Traffic Written: %s bytes\n", readTraffic, writtenTraffic)
		if _, err := conn.Write([]byte(response)); err != nil {
			handler.Logger.Printf("[ERROR] Failed to send traffic response: %v", err)
			return err
		}
		handler.Logger.Printf("[INFO] Traffic Read: %s bytes, Traffic Written: %s bytes", readTraffic, writtenTraffic)
		return nil
	case "stop":
		if err := handler.ClearIptablesIPv4(); err != nil {
			handler.Logger.Printf("[ERROR] %v", err)
			_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
			return err
		}
		if err := handler.StopPrivoxyService(); err != nil {
			handler.Logger.Printf("[ERROR] %v", err)
			_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
			return err
		}
		if err := handler.StopTorService(); err != nil {
			handler.Logger.Printf("[ERROR] %v", err)
			_, _ = conn.Write([]byte(fmt.Sprintf("Error: %v\n", err)))
			return err
		}
		_, _ = conn.Write([]byte("Done\n"))
		handler.Logger.Println("[INFO] Tor service stopped successfully.")
		return nil
	default:
		msg := fmt.Sprintf("Unknown command: %s\nAvailable commands: start, switch, traffic, stop\n", command)
		handler.Logger.Println("[WARN] " + msg)
		_, _ = conn.Write([]byte(msg))
		return nil
	}
}
