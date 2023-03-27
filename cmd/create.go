package cmd

import (
	"gomtastsflare/resource"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new resource",
	Long:  `Create a new resource`,
	RunE:  resource.ResourceCreate,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(resource.ResourceCreateCmd)
	createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
