package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jkt48-showroom-cli",
	Short: "JKT48 Showroom CLI is a CLI tool to get information about JKT48 members and their live stream",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}