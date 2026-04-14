package storage

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

const defaultDataFile = "bin/data/daydaynote.jsonl"

// Store persists notes in JSONL format, one JSON object per line.
type Store struct {
	path string
	mu   sync.Mutex
}

// NewStore creates a store backed by the given file path.
func NewStore(path string) *Store {
	return &Store{path: path}
}

// DefaultStore returns the repository-local JSONL store.
func DefaultStore() *Store {
	return NewStore(defaultDataFile)
}

// Path returns the underlying file path.
func (s *Store) Path() string {
	return s.path
}

// Append adds a note as a single JSONL row.
func (s *Store) Append(note Note) error {
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
	return enc.Encode(note)
}

// List reads every note from the JSONL file.
func (s *Store) List() ([]Note, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	f, err := os.Open(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Note{}, nil
		}
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	notes := make([]Note, 0)
	for scanner.Scan() {
		var note Note
		if err := json.Unmarshal(scanner.Bytes(), &note); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return notes, nil
}

// Replace rewrites the entire JSONL file with the provided notes.
func (s *Store) Replace(notes []Note) error {
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
	for _, note := range notes {
		if err := enc.Encode(note); err != nil {
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

func ensureParentDir(path string) error {
	dir := filepath.Dir(path)
	if dir == "." || dir == "" {
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}
