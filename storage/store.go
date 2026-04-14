package storage

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 负责把记录持久化到 JSONL 文件中，每行保存一条 JSON 记录
type Store struct {
	path string
	mu   sync.Mutex
}

// 根据指定文件路径创建一个存储实例
func NewStore(path string) *Store {
	return &Store{path: path}
}

// 返回默认存储，文件位于 exe 同级的 data 目录下
func DefaultStore() *Store {
	return NewStore(defaultDataPath())
}

// 返回当前存储使用的文件路径
func (s *Store) Path() string {
	return s.path
}

// 将一条记录追加写入 JSONL 文件末尾
func (s *Store) Append(record Record) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := ensureParentDir(s.path); err != nil {
		return err
	}

	f, err := os.OpenFile(s.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	return enc.Encode(record)
}

// 读取 JSONL 文件中的所有记录
func (s *Store) List() ([]Record, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	f, err := os.Open(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Record{}, nil
		}
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	records := make([]Record, 0)
	for scanner.Scan() {
		var record Record
		if err := json.Unmarshal(scanner.Bytes(), &record); err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return records, nil
}

// 用给定的记录列表整体重写 JSONL 文件
func (s *Store) Replace(records []Record) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := ensureParentDir(s.path); err != nil {
		return err
	}

	tmpPath := s.path + ".tmp"
	f, err := os.OpenFile(tmpPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(f)
	for _, record := range records {
		if err := enc.Encode(record); err != nil {
			f.Close()
			_ = os.Remove(tmpPath)
			return err
		}
	}

	if err := f.Close(); err != nil {
		_ = os.Remove(tmpPath)
		return err
	}
	return os.Rename(tmpPath, s.path)
}

// 确保目标文件所在的父目录存在
func ensureParentDir(path string) error {
	dir := filepath.Dir(path)
	if dir == "." || dir == "" {
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}

// 返回 exe 同级的 data 目录；如果无法获取 exe，就退回到当前目录下的 data
func dataDir() string {
	exe, err := os.Executable()
	if err != nil {
		return filepath.Clean("data")
	}

	return filepath.Join(filepath.Dir(exe), "data")
}

func defaultDataPath() string {
	now := time.Now()
	return filepath.Join(
		dataDir(),
		now.Format("2006"),
		now.Format("01")+".jsonl",
	)
}
