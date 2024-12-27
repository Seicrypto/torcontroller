package initializer

import (
	"embed"
	"errors"
	"os"

	"github.com/Seicrypto/torcontroller/internal/controller"
)

// TemplateProvider abstracts template file access.
type TemplateProvider interface {
	ReadFile(name string) ([]byte, error)
}

// EmbedFSWrapper wraps embed.FS to implement TemplateProvider.
type EmbedFSWrapper struct {
	FS embed.FS
}

func (w *EmbedFSWrapper) ReadFile(name string) ([]byte, error) {
	return w.FS.ReadFile(name)
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

type FileSystem interface {
	Stat(name string) (os.FileInfo, error)
	ReadFile(name string) ([]byte, error)
	MkdirAll(path string, perm os.FileMode) error
}

type RealFileSystem struct{}

func (r *RealFileSystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (r *RealFileSystem) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (r *RealFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// Initializer is responsible for system and config validations.
type Initializer struct {
	Templates     TemplateProvider
	CommandRunner controller.CommandRunner
	FileSystem    FileSystem
}

// NewInitializer creates a new Initializer with a given CommandRunner and Templates.
func NewInitializer(templates TemplateProvider, runner controller.CommandRunner, fs FileSystem) *Initializer {
	return &Initializer{
		Templates:     templates,
		CommandRunner: runner,
		FileSystem:    fs,
	}
}
