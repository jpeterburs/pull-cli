package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current version of pull-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull-cli version 0.1.0")
	},
}
