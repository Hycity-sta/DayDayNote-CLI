package cmd

import (
	"fmt"

	"daydaynote/i18n"

	"github.com/spf13/cobra"
)

const version = "0.1"

// version 命令
func Version() *cobra.Command {
	return &cobra.Command{
		Use:   i18n.T("cmd.version.use"),
		Short: i18n.T("cmd.version.short"),
		Run:   VersionHandle,
	}
}

// 输出当前版本信息
func VersionHandle(cmd *cobra.Command, args []string) {
	fmt.Printf(i18n.T("msg.version"), version)
}
