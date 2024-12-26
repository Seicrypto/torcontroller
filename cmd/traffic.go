package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var TrafficCmd = &cobra.Command{
	Use:   "traffic",
	Short: "Check the current traffic metrics for the Tor service",
	Run: func(cmd *cobra.Command, args []string) {
		handler := &SocketInteractionHandler{
			Adapter: &UnixSocketAdapter{SocketPath: socketPath},
		}

		response, err := handler.SendCommandAndGetResponse("traffic")
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}

		fmt.Printf("Response: %s\n", response)
	},
}
