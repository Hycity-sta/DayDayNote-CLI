package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "daydaynote",
	Short: "daydaynote is a small CLI app built with Cobra",
	Long:  "daydaynote is a small CLI app built with Cobra. Put each command in cmd/ for easy maintenance.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(helloCmd)
}
