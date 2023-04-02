package cmd

import (
	"gomtastsflare/resource"

	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify DNS Records and Web server Configuration",
	Long:  `Verify DNS Records and Web server Configuration`,
	RunE:  resource.ResourceVerify,
}

func init() {
	rootCmd.AddCommand(verifyCmd)
	verifyCmd.Flags().StringP("domain", "d", "", "Domain to Update (Required)")
	verifyCmd.MarkFlagRequired("domain")
}
