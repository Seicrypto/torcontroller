package cmd_test

import (
	"net"
)

// MockSocket is a mock implementation of ConnectionAdapter.
type MockSocket struct {
	ResponseMap map[string]string // Maps written commands to responses
	CloseSignal chan struct{}     // Signal to stop the goroutine
	Error       error             // Simulated connection error
}

// Dial simulates establishing a connection and returns a mock net.Conn.
func (m *MockSocket) Dial() (net.Conn, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	client, server := net.Pipe() // Simulate a connection
	go func() {
		defer server.Close()

		buf := make([]byte, 1024)
		for {
			select {
			case <-m.CloseSignal:
				return // Exit the loop when the close signal is received
			default:
				n, err := server.Read(buf)
				if err != nil {
					return // Connection closed
				}

				if response, exists := m.ResponseMap[string(buf[:n])]; exists {
					server.Write([]byte(response)) // Write the corresponding response
				} else {
					server.Write([]byte("5XX Command not recognized\n"))
				}
			}
		}
	}()

	return client, nil
}

// MockConnectionAdapter emulates ConnectionAdapter.
type MockConnectionAdapter struct {
	MockConn *MockSocket
	FailDial bool
}
