package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SwitchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch the Tor IP",
	Run: func(cmd *cobra.Command, args []string) {
		handler := &SocketInteractionHandler{
			Adapter: &UnixSocketAdapter{SocketPath: socketPath},
		}

		response, err := handler.SendCommandAndGetResponse("switch")
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}

		fmt.Printf("Response: %s\n", response)
	},
}
