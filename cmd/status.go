package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the current state of the Torcontroller listener",
	Run: func(cmd *cobra.Command, args []string) {
		handler := &SocketInteractionHandler{
			Adapter: &UnixSocketAdapter{SocketPath: socketPath},
		}

		response, err := handler.SendCommand("status")
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}

		fmt.Printf("Response: %s\n", response)
	},
}
