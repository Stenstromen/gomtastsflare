package cmd

import (
	"gomtastsflare/resource"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update DNS Records and/or Nginx Configuration",
	Long:  `Update DNS Records and/or Nginx Configuration`,
	RunE:  resource.ResourceUpdate,
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("domain", "d", "", "Domain to Update (Required)")
	updateCmd.Flags().StringP("mx", "m", "", "MX Record(s) to Mailserver (ex mx1.com,mx2.com)")
	updateCmd.Flags().StringP("ipv4", "4", "", "IPv4 Address to Webserver")
	updateCmd.Flags().StringP("ipv6", "6", "", "IPv6 Address to Webserver")
	updateCmd.Flags().StringP("rua", "r", "", "Report Email Address for MTA-STS")
	updateCmd.MarkFlagRequired("domain")
}
