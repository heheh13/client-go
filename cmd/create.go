package cmd

import (
	"fmt"

	"github.com/heheh13/client-go/api"

	"github.com/spf13/cobra"
)

var (
	// wanted to use for  updating container images
	//looks like it will need kwargs for that
	// seems ugly to edit parameters list every time a new things added in the list
	image string
	///responsible for creating deployment resources
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "creating a api resource",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
			api.CreateDelployment()
		},
	}
	//updating
	// planned to use some flags
	// also planing to read from files
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "update api resources",
		Run: func(cmd *cobra.Command, args []string) {
			api.UpdateDeployment()
		},
	}
	// Get command
	// list all the deployment resources
	// planing for multiple resources geting using args?

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get api resources",
		Run: func(cmd *cobra.Command, args []string) {
			api.GetDeployment()
		},
	}
	//Delete api resources
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete api resources",
		Run: func(cmd *cobra.Command, args []string) {
			api.DeleteDeployment()
		},
	}
)

func init() {
	updateCmd.PersistentFlags().StringVarP(&image, "image", "i", "nginx:1.13", "container Image")
	rootCmd.AddCommand(createCmd, updateCmd, getCmd, deleteCmd)

}
