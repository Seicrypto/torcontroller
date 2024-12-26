package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Seicrypto/torcontroller/internal/singleton/logger"
	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:   "stop [socketPath]",
	Short: "Stop a Torcontroller listener",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logger.GetLogger()

		// Send the "stop" command to the listener via the socket
		handler := &SocketInteractionHandler{
			Adapter: &UnixSocketAdapter{SocketPath: socketPath},
		}

		response, err := handler.SendCommandAndGetResponse("stop")
		if err != nil {
			logger.Error(fmt.Sprintf("Error sending command: %v", err))
			fmt.Printf("Error sending command: %v\n", err)
			return
		}

		if response != "done\n" {
			logger.Warn(fmt.Sprintf("Unexpected response from server: %s", response))
			fmt.Printf("Unexpected response from server: %s\n", response)
			return
		}

		logger.Info("Server confirmed successful stop.")
		fmt.Println("Server confirmed successful stop.")

		// Read the PID file
		data, err := os.ReadFile(pidFile)
		if err != nil {
			fmt.Printf("Error reading PID file: %v\n", err)
			return
		}

		pid, err := strconv.Atoi(string(data))
		if err != nil {
			fmt.Printf("Error parsing PID: %v\n", err)
			return
		}

		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Printf("Error finding process: %v\n", err)
			return
		}

		err = proc.Kill()
		if err != nil {
			fmt.Printf("Error killing process: %v\n", err)
			return
		}

		os.Remove(socketPath)
		os.Remove(pidFile)
		logger.Info(fmt.Sprintf("Torcontroller listener at %s stopped successfully.\n", socketPath))
		fmt.Printf("Torcontroller listener at %s stopped successfully.\n", socketPath)
	},
}
