package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the app version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("daydaynote %s\n", version)
		},
	}
}
