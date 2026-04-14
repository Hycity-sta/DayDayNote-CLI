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
		Short: i18n.T(i18n.RootShort),
		Long:  i18n.T(i18n.RootLong),
	}

	// 挂载子命令
	rootCmd.AddCommand(cmd.Add())
	rootCmd.AddCommand(cmd.List())
	rootCmd.AddCommand(cmd.Delete())
	rootCmd.AddCommand(cmd.Edit())
	rootCmd.AddCommand(cmd.Lang())
	rootCmd.AddCommand(cmd.Version())
	cmd.ConfigureHelp(rootCmd)

	// 处理错误
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
