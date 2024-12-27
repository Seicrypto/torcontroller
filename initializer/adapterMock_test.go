package initializer_test

import (
	"errors"
	"os"
	"syscall"
	"time"
)

// MockCommandRunner is a mock implementation of the CommandRunner interface for testing.
type MockCommandRunner struct {
	Output string
	Error  error
}

func (m *MockCommandRunner) Run(name string, args ...string) (string, error) {
	if m.Error != nil {
		return "", m.Error
	}
	return m.Output, nil
}

// MockTemplates is a mock implementation of TemplateProvider for testing.
type MockTemplates struct {
	Files map[string][]byte
	Error error
}

func (m *MockTemplates) ReadFile(name string) ([]byte, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	content, exists := m.Files[name]
	if !exists {
		return nil, errors.New("file not found")
	}
	return content, nil
}

// MockFileInfo is a mock implementation of os.FileInfo for testing.
type MockFileInfo struct {
	content []byte
	mode    os.FileMode
	uid     uint32
	gid     uint32
}

func (m *MockFileInfo) Name() string       { return "mockfile" }
func (m *MockFileInfo) Size() int64        { return 0 }
func (m *MockFileInfo) Mode() os.FileMode  { return m.mode }
func (m *MockFileInfo) ModTime() time.Time { return time.Now() }
func (m *MockFileInfo) IsDir() bool        { return false }
func (m *MockFileInfo) Sys() interface{} {
	return &syscall.Stat_t{
		Uid: m.uid,
		Gid: m.gid,
	}
}

// MockFileSystem is a mock implementation of FileSystem for testing.
type MockFileSystem struct {
	Files       map[string]*MockFileInfo
	Error       error
	MkdirErrors map[string]error
}

func (m *MockFileSystem) Stat(name string) (os.FileInfo, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	info, exists := m.Files[name]
	if !exists {
		return nil, errors.New("file not found")
	}
	return info, nil
}

func (m *MockFileSystem) ReadFile(name string) ([]byte, error) {
	info, exists := m.Files[name]
	if !exists {
		return nil, errors.New("file not found")
	}
	return info.content, nil
}

func (m *MockFileSystem) MkdirAll(path string, perm os.FileMode) error {
	if err, exists := m.MkdirErrors[path]; exists {
		return err
	}
	// Simulate creating a directory
	m.Files[path] = &MockFileInfo{mode: perm}
	return nil
}
