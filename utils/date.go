package utils

// 统一放命令层共用的日期和索引处理，避免 list、delete、edit 重复写同样逻辑
import (
	"fmt"
	"strconv"
	"time"

	"daydaynote/i18n"
	"daydaynote/storage"
)

// 统一处理命令里的日期参数，未传时默认使用今天
func ResolveDate(value string) (time.Time, error) {
	if value == "" {
		return time.Now(), nil
	}

	date, err := storage.ParseDate(value)
	if err != nil {
		return time.Time{}, fmt.Errorf(i18n.T(i18n.ErrDateInvalid), value)
	}

	return date, nil
}

// 判断两个时间是否落在同一个本地自然日
func SameDate(a, b time.Time) bool {
	a = a.Local()
	b = b.Local()
	return a.Year() == b.Year() && a.Month() == b.Month() && a.Day() == b.Day()
}

// 把命令行输入的索引解析成从 1 开始的正整数
func ParsePositiveIndex(value string) (int, error) {
	index, err := strconv.Atoi(value)
	if err != nil || index <= 0 {
		return 0, fmt.Errorf(i18n.T(i18n.ErrIndexInvalid), value)
	}

	return index, nil
}
