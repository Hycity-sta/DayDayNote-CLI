package storage

import "time"

// Note 表示一条便签数据，会被持久化到 JSONL 文件中。
type Note struct {
	// ID 是便签的唯一标识。
	ID string `json:"id"`
	// Title 是便签标题。
	Title string `json:"title"`
	// Content 是便签正文内容。
	Content string `json:"content"`
	// CreatedAt 是便签创建时间。
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt 是便签最后更新时间。
	UpdatedAt time.Time `json:"updated_at"`
}
