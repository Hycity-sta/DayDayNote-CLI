package daydaynote_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"daydaynote/i18n"
)

// withLang 会临时切换当前语言，并在测试结束后恢复原来的语言。
func withLang(t *testing.T, lang string) {
	t.Helper()

	old := i18n.Lang()
	i18n.SetLang(lang)
	t.Cleanup(func() {
		i18n.SetLang(old)
	})
}

// captureStdout 执行函数，并返回它写到标准输出的内容。
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("create pipe: %v", err)
	}

	os.Stdout = w
	defer func() {
		os.Stdout = oldStdout
	}()

	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("close pipe writer: %v", err)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("read stdout: %v", err)
	}
	if err := r.Close(); err != nil {
		t.Fatalf("close pipe reader: %v", err)
	}

	return buf.String()
}
