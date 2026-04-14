package storage

import "time"

// 一条记录数据，会被持久化到 JSONL 文件中
type Record struct {
	// 记录的唯一标识
	ID string `json:"id"`
	// 记录标题
	Title string `json:"title"`
	// 记录正文内容
	Content string `json:"content"`
	// 记录创建时间
	CreatedAt time.Time `json:"created_at"`
	// 记录最后更新时间
	UpdatedAt time.Time `json:"updated_at"`
}
