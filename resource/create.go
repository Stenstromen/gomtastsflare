package resource

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ResourceCreateCmd = &cobra.Command{
	RunE: ResourceCreate,
}

func ResourceCreate(cmd *cobra.Command, args []string) error {
	fmt.Println("Create Cloudflare Records here")
	return nil
}
