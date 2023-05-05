package cmd

import (
	"gomtastsflare/resource"

	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Read report file and print to stdout",
	Long:  `Read report file and print to stdout`,
	RunE:  resource.ResourceReport,
}

func init() {
	rootCmd.AddCommand(reportCmd)
	reportCmd.Flags().StringP("file", "f", "", "Reportfile to read (Required)")
	reportCmd.MarkFlagRequired("file")
}
