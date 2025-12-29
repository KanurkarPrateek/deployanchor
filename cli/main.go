package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "anchor",
	Short: "DeployAnchor CLI",
	Long:  `A simple CLI to interact with the DeployAnchor PaaS.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to DeployAnchor CLI!")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
