package cmd

import (
	"daydaynote/i18n"

	"github.com/spf13/cobra"
)

// help 命令
func Help() *cobra.Command {
	return &cobra.Command{
		Use:   i18n.T(i18n.CmdHelpUse),
		Short: i18n.T(i18n.CmdHelpShort),
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			// 返回默认的help说明
			return cmd.Root().Help()
		},
	}
}
