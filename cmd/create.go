package cmd

import (
	"gomtastsflare/resource"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create required DNS records and Nginx configuration",
	Long:  `Create required DNS records and Nginx configuration`,
	RunE:  resource.ResourceCreate,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("domain", "d", "", "Domain to Update or Create (Required)")
	createCmd.Flags().StringP("mx", "m", "", "MX Record to Mailserver (Required)")
	createCmd.Flags().StringP("ipv4", "4", "", "IPv4 Address to Webserver (Required)")
	createCmd.Flags().StringP("ipv6", "6", "", "IPv6 Address to Webserver")
	createCmd.Flags().StringP("rua", "r", "", "Report Email Address for MTA-STS (Required)")
	createCmd.MarkFlagRequired("domain")
	createCmd.MarkFlagRequired("ipv4")
	createCmd.MarkFlagRequired("mx")
	createCmd.MarkFlagRequired("rua")

}
