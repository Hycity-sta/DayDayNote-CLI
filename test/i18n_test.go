package daydaynote_test

import (
	"testing"

	"daydaynote/config"
	"daydaynote/i18n"
)

// 验证语言切换以及缺失 key 时的回退行为。
func TestSetLangAndFallback(t *testing.T) {
	withLang(t, config.LanguageEnglish)

	i18n.SetLang(config.LanguageChinese)
	if got := i18n.Lang(); got != config.LanguageChinese {
		t.Fatalf("Lang() = %q，期望 %q", got, config.LanguageChinese)
	}
	if got := i18n.T(i18n.CmdHelpShort); got != "显示命令帮助信息" {
		t.Fatalf("T(cmd.help.short) = %q，期望 %q", got, "显示命令帮助信息")
	}

	i18n.SetLang("fr")
	if got := i18n.Lang(); got != config.LanguageEnglish {
		t.Fatalf("不支持的语言切换后 Lang() = %q，期望 %q", got, config.LanguageEnglish)
	}
	if got := i18n.T("missing.key"); got != "missing.key" {
		t.Fatalf("缺失 key 的回退结果 = %q，期望 key 本身", got)
	}
}
