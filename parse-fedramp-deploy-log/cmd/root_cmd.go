package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "app",
}

func init() {
	rootCmd.AddCommand(unzipCMD)
}

func Execute() error {
	err := rootCmd.Execute()
	return err
}
