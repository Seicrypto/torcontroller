package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/Seicrypto/torcontroller/internal/controller"
	"github.com/Seicrypto/torcontroller/internal/services/logger"
	"github.com/spf13/cobra"
)

var StartBackgroundCmd = &cobra.Command{
	Use:   "start-background",
	Short: "Start Torcontroller listener as a background process",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logger.GetLogger()
		logger.Info(fmt.Sprintf("Listener started successfully at %s.\n", socketPath))
		os.Remove(socketPath)
		listener, err := net.Listen("unix", socketPath)
		if err != nil {
			logger.Error(err.Error())
			return
		}
		defer func() {
			listener.Close()
			os.Remove(socketPath)
		}()
		os.Chmod(socketPath, 0777)

		// Loop for accepting connections
		for {
			logger.Info("Waiting for connection...")
			conn, err := listener.Accept()
			if err != nil {
				logger.Warn(fmt.Sprintf("Error accepting connection: %v\n", err))
				continue
			}
			logger.Info("Connection established")

			go controller.HandleConnection(conn, socketPath, listener)
		}
	},
}
