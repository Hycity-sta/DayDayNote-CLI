package daydaynote_test

import (
	"strings"
	"testing"

	"daydaynote/cmd"
	"daydaynote/config"
	"daydaynote/i18n"

	"github.com/spf13/cobra"
)

// 验证自定义 help 命令的元信息和无参数行为。
func TestHelpCommand(t *testing.T) {
	withLang(t, config.LanguageChinese)

	command := cmd.Help()
	if got := command.Use; got != "help" {
		t.Fatalf("Help().Use = %q，期望 %q", got, "help")
	}
	if got := command.Short; got != "显示命令帮助信息" {
		t.Fatalf("Help().Short = %q，期望 %q", got, "显示命令帮助信息")
	}

	if err := command.Args(&cobra.Command{}, nil); err != nil {
		t.Fatalf("Help().Args() 不应拒绝空参数：%v", err)
	}
	if err := command.Args(&cobra.Command{}, []string{"delete"}); err == nil {
		t.Fatal("Help().Args() 应该拒绝任意位置参数")
	}

	root := &cobra.Command{Use: "daydaynote"}
	root.AddCommand(command)

	output := captureStdout(t, func() {
		root.SetArgs([]string{"help"})
		if err := root.Execute(); err != nil {
			t.Fatalf("执行 help 命令出错：%v", err)
		}
	})
	if !strings.Contains(output, "daydaynote") {
		t.Fatalf("help 输出 %q 不包含根命令名称", output)
	}
}

// 锁定 version 命令的输出格式。
func TestVersionHandle(t *testing.T) {
	withLang(t, config.LanguageEnglish)

	output := captureStdout(t, func() {
		cmd.VersionHandle(&cobra.Command{}, nil)
	})
	if got := strings.TrimSpace(output); got != "daydaynote 0.1" {
		t.Fatalf("VersionHandle() 输出 = %q，期望 %q", got, "daydaynote 0.1")
	}
}

// 确保命令名称和简介仍然来自当前语言环境。
func TestCommandStringsFollowLocale(t *testing.T) {
	withLang(t, config.LanguageChinese)

	if got := cmd.Version().Short; got != i18n.T(i18n.CmdVersionShort) {
		t.Fatalf("Version().Short = %q，期望 %q", got, i18n.T(i18n.CmdVersionShort))
	}
}
