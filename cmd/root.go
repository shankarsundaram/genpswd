/*
Copyright © 2023 Shankar Sundaram shankar.sundaram@outlook.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "genpswd",
	Short: "Generate Random Password",
	Long:  `Generate Random Password based of customizable options.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
