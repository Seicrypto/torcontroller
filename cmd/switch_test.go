package cmd_test

import (
	"net"
	"testing"

	"github.com/Seicrypto/torcontroller/cmd"
)

func TestSwitchCmd(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	go func() {
		buf := make([]byte, 1024)
		n, _ := server.Read(buf)
		received := string(buf[:n])
		if received != "switch" {
			t.Errorf("Expected command 'switch', but got '%s'", received)
		}
		server.Write([]byte("ACK\n"))
	}()

	handler := &cmd.SocketInteractionHandler{
		Adapter: &MockAdapter{Client: client, Server: server},
	}

	response, err := handler.SendCommandAndGetResponse("switch")
	if err != nil {
		t.Fatalf("Execution failed: %v", err)
	}

	expected := "ACK\n"
	if response != expected {
		t.Errorf("Expected response '%s', but got '%s'", expected, response)
	}
}
