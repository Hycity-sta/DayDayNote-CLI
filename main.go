package main

import (
	"fmt"
	"os"

	"daydaynote/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "daydaynote",
	Short: "daydaynote is a small CLI app built with Cobra",
	Long:  "daydaynote is a small CLI app built with Cobra. Put each command in cmd/ for easy maintenance.",
}

func main() {
	rootCmd.AddCommand(cmd.NewVersionCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
