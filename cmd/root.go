package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "...",
	Short:   "rootCommand",
	Version: "v1.1.1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create <resource> name")
		fmt.Println("get <resource> name")
		fmt.Println("delete <resource> name")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
