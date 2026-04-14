package storage

import "path/filepath"

// 把 filepath 能力收口到这里，便于复用路径规则
func filepathDir(path string) string {
	return filepath.Dir(path)
}

// 把多个路径片段拼成一个规范路径
func filepathJoin(elem ...string) string {
	return filepath.Join(elem...)
}

// 清理路径中的冗余片段
func filepathClean(path string) string {
	return filepath.Clean(path)
}
