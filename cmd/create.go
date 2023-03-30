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
	createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringP("name", "n", "", "Resource Name")
	createCmd.Flags().StringP("username", "u", "", "Resource Username")
	createCmd.Flags().String("uri", "", "Resource URI")
	createCmd.Flags().StringP("password", "p", "", "Resource Password")
	createCmd.Flags().StringP("description", "d", "", "Resource Description")
	createCmd.Flags().StringP("folderParentID", "f", "", "Folder in which to create the Resource")

	createCmd.MarkFlagRequired("name")

}
