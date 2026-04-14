package daydaynote_test

import (
	"strings"
	"testing"
	"time"

	"daydaynote/config"
	"daydaynote/utils"

	"github.com/spf13/cobra"
)

// 覆盖通用的 Cobra 参数校验器，包含成功和失败两种情况。
func TestArgValidators(t *testing.T) {
	withLang(t, config.LanguageChinese)

	noArgs := utils.NoArgs()
	if err := noArgs(&cobra.Command{}, nil); err != nil {
		t.Fatalf("NoArgs() 不应拒绝空参数：%v", err)
	}
	if err := noArgs(&cobra.Command{}, []string{"x"}); err == nil || !strings.Contains(err.Error(), "不接受位置参数") {
		t.Fatalf("NoArgs() 错误信息不符合预期：%v", err)
	}

	exact := utils.ExactArgs(2)
	if err := exact(&cobra.Command{}, []string{"a", "b"}); err != nil {
		t.Fatalf("ExactArgs() 不应拒绝正确数量的参数：%v", err)
	}
	if err := exact(&cobra.Command{}, []string{"a"}); err == nil || !strings.Contains(err.Error(), "需要 2 个参数") {
		t.Fatalf("ExactArgs() 错误信息不符合预期：%v", err)
	}

	minimum := utils.MinimumNArgs(2)
	if err := minimum(&cobra.Command{}, []string{"a", "b"}); err != nil {
		t.Fatalf("MinimumNArgs() 不应拒绝足够数量的参数：%v", err)
	}
	if err := minimum(&cobra.Command{}, []string{"a"}); err == nil || !strings.Contains(err.Error(), "至少需要 2 个参数") {
		t.Fatalf("MinimumNArgs() 错误信息不符合预期：%v", err)
	}
}

// 确保正整数索引解析会拒绝非正数和非法输入。
func TestParsePositiveIndex(t *testing.T) {
	withLang(t, config.LanguageChinese)

	index, err := utils.ParsePositiveIndex("3")
	if err != nil {
		t.Fatalf("ParsePositiveIndex(3) 出错：%v", err)
	}
	if index != 3 {
		t.Fatalf("ParsePositiveIndex(3) = %d，期望 3", index)
	}

	if _, err := utils.ParsePositiveIndex("0"); err == nil || !strings.Contains(err.Error(), "无效索引") {
		t.Fatalf("ParsePositiveIndex(0) 错误信息不符合预期：%v", err)
	}
}

// 验证命令层会用到的日期比较和日期解析工具。
func TestSameDateAndResolveDate(t *testing.T) {
	withLang(t, config.LanguageEnglish)

	dayA := time.Date(2026, 4, 14, 9, 0, 0, 0, time.Local)
	dayB := time.Date(2026, 4, 14, 23, 59, 59, 0, time.Local)
	dayC := time.Date(2026, 4, 15, 0, 0, 0, 0, time.Local)

	if !utils.SameDate(dayA, dayB) {
		t.Fatal("SameDate() 应该把同一天视为相等")
	}
	if utils.SameDate(dayA, dayC) {
		t.Fatal("SameDate() 应该把不同的自然日视为不相等")
	}

	got, err := utils.ResolveDate("2026-4-14")
	if err != nil {
		t.Fatalf("ResolveDate() 出错：%v", err)
	}
	if got.Year() != 2026 || got.Month() != 4 || got.Day() != 14 {
		t.Fatalf("ResolveDate() = %v，期望 2026-04-14", got)
	}

	if _, err := utils.ResolveDate("not-a-date"); err == nil || !strings.Contains(err.Error(), "invalid date") {
		t.Fatalf("ResolveDate() 对非法日期的错误信息不符合预期：%v", err)
	}
}
