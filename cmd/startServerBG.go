package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/Seicrypto/torcontroller/internal/services/logger"
	"github.com/spf13/cobra"
)

var startBackgroundCmd = &cobra.Command{
	Use:   "start-background",
	Short: "Start Torcontroller listener as a background process",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		logger, err := logger.CreateLogger()
		if err != nil {
			fmt.Println("Error creating logger:", err)
			return
		}
		logger.Printf("Listener started successfully at %s.\n", socketPath)
		os.Remove(socketPath)
		listener, err := net.Listen("unix", socketPath)
		if err != nil {
			logger.Fatal(err)
			return
		}
		defer func() {
			listener.Close()
			os.Remove(socketPath)
		}()
		os.Chmod(socketPath, 0777)

		// Loop for accepting connections
		for {
			logger.Println("Waiting for connection...")
			conn, err := listener.Accept()
			if err != nil {
				logger.Printf("Error accepting connection: %v\n", err)
				continue
			}
			logger.Println("Connection established")

			go func(c net.Conn) {
				defer c.Close()
				buf := make([]byte, 1024)
				n, err := c.Read(buf)
				if err != nil {
					logger.Printf("Error reading from connection: %v\n", err)
					return
				}

				logger.Printf("Received: %s\n", string(buf[:n]))

				_, err = c.Write([]byte("ACK\n"))
				if err != nil {
					logger.Printf("Error writing to connection: %v\n", err)
				}
			}(conn)
		}
	},
}
