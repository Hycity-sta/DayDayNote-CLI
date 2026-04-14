package cmd

import (
	"fmt"

	"daydaynote/config"
	"daydaynote/i18n"

	"github.com/spf13/cobra"
)

// lang 命令
func Lang() *cobra.Command {
	return &cobra.Command{
		Use:   i18n.T("cmd.lang.use"),
		Short: i18n.T("cmd.lang.short"),
		Args:  cobra.ExactArgs(1),
		RunE:  LangHandle,
	}
}

// 设置默认语言并写入配置文件
func LangHandle(cmd *cobra.Command, args []string) error {
	// 传递第一个参数
	return ApplyLanguage(args[0])
}

// 切换当前语言，并把语言配置写入 exe 同级目录
func ApplyLanguage(lang string) error {
	// 先校验传入的语言是否在支持范围内。
	if !config.IsSupportedLanguage(lang) {
		return fmt.Errorf(i18n.T("err.lang.unsupported"), lang)
	}

	// 先更新当前进程里的语言，保证后续提示信息立即生效。
	i18n.SetLang(lang)

	// 再把语言写入配置文件，让下次启动时也能继续使用这个设置。
	if err := config.SaveLanguage(lang); err != nil {
		return err
	}

	// 配置写入成功后，输出设置完成提示。
	fmt.Printf(i18n.T("msg.lang.set"), lang)
	return nil
}
