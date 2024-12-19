package cmd_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/Seicrypto/torcontroller/cmd"
)

func TestStatusCmd(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	go func() {
		buf := make([]byte, 1024)
		n, _ := server.Read(buf)
		fmt.Printf("Read %d bytes: %s\n", n, string(buf[:n]))
		server.Write([]byte("ACK\n"))
	}()

	handler := &cmd.SocketInteractionHandler{
		Adapter: &MockAdapter{Client: client, Server: server},
	}

	response, err := handler.SendCommand("status")
	if err != nil {
		t.Fatalf("Execution failed: %v", err)
	}

	expected := "ACK\n"
	if response != expected {
		t.Errorf("Expected response '%s', but got '%s'", expected, response)
	}
}
