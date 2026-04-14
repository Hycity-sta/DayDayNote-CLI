package storage

import (
	"fmt"
	"time"
)

var supportedDateLayouts = []string{
	"2006/1/2",
	"2006/01/02",
	"2006-1-2",
	"2006-01-02",
}

// 返回指定日期对应的存储，文件仍按“年/月.jsonl”组织
func StoreForDate(date time.Time) *Store {
	return NewStore(dataPathForDate(date))
}

// 按输入格式解析日期，只接受项目里约定的两种分隔符格式
func ParseDate(value string) (time.Time, error) {
	for _, layout := range supportedDateLayouts {
		if date, err := time.ParseInLocation(layout, value, time.Local); err == nil {
			return date, nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid date: %s", value)
}

// 默认存储路径按“data/年/月.jsonl”组织，
// 例如：data/2026/04.jsonl。
func dataPathForDate(date time.Time) string {
	return filepathJoin(
		dataDir(),
		date.Format("2006"),
		date.Format("01")+".jsonl",
	)
}
