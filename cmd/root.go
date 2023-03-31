package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gomtastsflare",
	Short: "Go binary for creating/updating MTA-STS records on Cloudflare, and create the accompanying Nginx configuration.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
