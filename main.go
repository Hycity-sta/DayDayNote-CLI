package main

import (
	"fmt"
	"os"

	"daydaynote/cmd"
	"daydaynote/config"
	"daydaynote/i18n"

	"github.com/spf13/cobra"
)

func main() {
	// 设置语言
	i18n.SetLang(config.LoadLanguage())

	// 实例化rootcmd，这个就是主程序
	rootCmd := &cobra.Command{
		Use:   "daydaynote",
		Short: i18n.T("root.short"),
		Long:  i18n.T("root.long"),
	}

	// 挂载子命令
	rootCmd.AddCommand(cmd.Lang())
	rootCmd.AddCommand(cmd.Version())

	// 处理错误
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
