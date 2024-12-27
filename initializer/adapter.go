package initializer

import (
	"embed"

	"github.com/Seicrypto/torcontroller/internal/controller"
)

// Initializer is responsible for system and config validations.
type Initializer struct {
	Templates     embed.FS
	CommandRunner controller.CommandRunner
}

// NewInitializer creates a new Initializer with a given CommandRunner and Templates.
func NewInitializer(templates embed.FS, runner controller.CommandRunner) *Initializer {
	return &Initializer{
		Templates:     templates,
		CommandRunner: runner,
	}
}
