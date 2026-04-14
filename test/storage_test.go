package daydaynote_test

import (
	"path/filepath"
	"strings"
	"testing"
	"time"

	"daydaynote/storage"
)

// 确认支持的日期格式都能解析成同一个自然日。
func TestParseDate(t *testing.T) {
	got, err := storage.ParseDate("2026/4/14")
	if err != nil {
		t.Fatalf("ParseDate() 出错：%v", err)
	}
	if got.Year() != 2026 || got.Month() != 4 || got.Day() != 14 {
		t.Fatalf("ParseDate() = %v，期望 2026-04-14", got)
	}

	got, err = storage.ParseDate("2026-04-14")
	if err != nil {
		t.Fatalf("ParseDate() 出错：%v", err)
	}
	if got.Year() != 2026 || got.Month() != 4 || got.Day() != 14 {
		t.Fatalf("ParseDate() = %v，期望 2026-04-14", got)
	}

	if _, err := storage.ParseDate("2026/14/4"); err == nil {
		t.Fatal("ParseDate() 应该拒绝非法日期")
	}
}

// 锁定按月分文件的存储路径规则。
func TestStoreForDatePath(t *testing.T) {
	date := time.Date(2026, 4, 14, 0, 0, 0, 0, time.Local)
	store := storage.StoreForDate(date)

	wantSuffix := filepath.Join("data", "2026", "04.jsonl")
	if !strings.HasSuffix(store.Path(), wantSuffix) {
		t.Fatalf("StoreForDate().Path() = %q，期望后缀 %q", store.Path(), wantSuffix)
	}
}
