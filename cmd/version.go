package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"--version", "-v"},
	Short:   "Show TorController version",
	Long:    "TorController CLI version 1.1.0\n\nVisit https://github.com/Seicrypto/torcontroller for release notes and updates.",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TorController version 1.1.0")
	},
}
