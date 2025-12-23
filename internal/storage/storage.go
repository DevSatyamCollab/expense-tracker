package storage

import (
	"encoding/json"
	"expense-tracker/internal/domain"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// --------------------
// IStorage
// --------------------
type IStorage interface {
	Save(tracker *domain.ExpenseTracker) error
	Load(tracker *domain.ExpenseTracker) error
}

const (
	dirPath       = ".expense-tracker/data/"
	fname         = "expenses.json"
	readwritePerm = 0600
	lock          = 0000
)

// ------------------
// storage
// ------------------
type JsonStorage struct {
	filename string
}

var (
	instance *JsonStorage
	once     sync.Once
)

// unsure storage
func unsureStorage() (string, error) {

	if err := os.MkdirAll(dirPath, 0700); err != nil {
		return "", fmt.Errorf("failed to create a directory: %w", err)
	}

	// custom filepath
	filePath := filepath.Join(dirPath, fname)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to create a file: %w", err)
		}

		if err := file.Close(); err != nil {
			return "", fmt.Errorf("failed to close the file: %w", err)
		}

		if err := os.Chmod(filePath, lock); err != nil {
			return "", fmt.Errorf("failed to lock the file: %w", err)
		}

	}

	return filePath, nil
}

// GetStorage
func GetStorage() (*JsonStorage, error) {
	var storageErr error

	once.Do(func() {
		file, err := unsureStorage()

		if err != nil {
			storageErr = fmt.Errorf("failed to initialize storage: %w", err)
			instance = nil
		} else {
			instance = &JsonStorage{filename: file}
		}
	})

	return instance, storageErr
}

// save method
func (s *JsonStorage) Save(tracker *domain.ExpenseTracker) error {
	if s.filename == "" {
		return domain.ErrEmptyStoragePath
	}

	// unlock the file
	if err := s.unlockFile(); err != nil {
		return err
	}

	// marshal
	jsonData, err := json.MarshalIndent(tracker, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	// writing the file
	if err = os.WriteFile(s.filename, jsonData, 0200); err != nil {
		return fmt.Errorf("failed to WriteFile: %w", err)
	}

	// lock the file
	if err = s.lockFile(); err != nil {
		return err
	}

	return nil
}

// load method
func (s *JsonStorage) Load(tracker *domain.ExpenseTracker) error {
	if s.filename == "" {
		return domain.ErrEmptyStoragePath
	}

	// unlock the file
	if err := s.unlockFile(); err != nil {
		return err
	}

	// reading the file
	content, err := os.ReadFile(s.filename)
	if err != nil {
		return fmt.Errorf("failed to read the file: %w", err)
	}

	if len(content) == 0 {
		return nil
	}

	// unmarshal the data
	if err = json.Unmarshal(content, tracker); err != nil {
		return fmt.Errorf("failed to unmarshal the data: %w", err)
	}

	// update the nextID
	tracker.NextID = tracker.UpdateNextID()

	// lock the file
	if err = s.lockFile(); err != nil {
		return err
	}

	return nil
}

// lock the file
func (s *JsonStorage) lockFile() error {
	if err := os.Chmod(s.filename, lock); err != nil {
		return fmt.Errorf("failed to lock the file: %w", err)
	}

	return nil
}

// unlock the file
func (s *JsonStorage) unlockFile() error {
	if err := os.Chmod(s.filename, readwritePerm); err != nil {
		return fmt.Errorf("failed to unlock the file: %w", err)
	}

	return nil
}
