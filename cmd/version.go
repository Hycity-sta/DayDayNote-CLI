package cmd

import (
	"fmt"

	"daydaynote/i18n"

	"github.com/spf13/cobra"
)

const version = "1.0"

// version 命令
func Version() *cobra.Command {
	return &cobra.Command{
		Use:   i18n.T(i18n.CmdVersionUse),
		Short: i18n.T(i18n.CmdVersionShort),
		Run:   VersionHandle,
	}
}

// 输出当前版本信息
func VersionHandle(cmd *cobra.Command, args []string) {
	fmt.Printf(i18n.T(i18n.MsgVersion), version)
}
