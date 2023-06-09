package gttp

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gttp",
	Short: "gttp - a command-line tool to ping HTTP servers",
	Long:  "ADD LONGER DESCRIPTION HERE",

	Run: func(cmd *cobra.Command, arg []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed while running gttp commands '%s'", err)
		os.Exit(1)
	}
}
