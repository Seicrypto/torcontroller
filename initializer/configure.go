package initializer

import (
	"embed"
	"fmt"
	"os"

	runneradapter "github.com/Seicrypto/torcontroller/internal/services/runnerAdapter"
)

//go:embed templates/*
var templates embed.FS

func PlaceTorServiceFile(runner runneradapter.CommandRunner) error {
	content, err := templates.ReadFile("templates/tor.service")
	if err != nil {
		return fmt.Errorf("failed to read tor service template: %w", err)
	}
	return writeServiceFile("/etc/systemd/system/tor.service", content, runner)
}

func PlacePrivoxyServiceFile(runner runneradapter.CommandRunner) error {
	content, err := templates.ReadFile("templates/privoxy.service")
	if err != nil {
		return fmt.Errorf("failed to read privoxy service template: %w", err)
	}
	return writeServiceFile("/etc/systemd/system/privoxy.service", content, runner)
}

func writeServiceFile(path string, content []byte, runner runneradapter.CommandRunner) error {
	tmpFile := "/tmp/service.tmp"
	if err := os.WriteFile(tmpFile, content, 0644); err != nil {
		return fmt.Errorf("failed to write temp file: %w", err)
	}

	cmd := []string{"sudo", "mv", tmpFile, path}
	if _, err := runner.Run(cmd[0], cmd[1:]...); err != nil {
		return fmt.Errorf("failed to move service file: %w", err)
	}
	return nil
}

func PlaceSudoersFile(runner runneradapter.CommandRunner) error {
	sudoersPath := "/etc/sudoers.d/torcontroller"

	content, err := templates.ReadFile("templates/sudoers.d/torcontroller")
	if err != nil {
		return fmt.Errorf("failed to read sudoers template: %w", err)
	}

	tmpFile, err := os.CreateTemp("", "torcontroller-sudoers-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temporary sudoers file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(content); err != nil {
		return fmt.Errorf("failed to write to temporary sudoers file: %w", err)
	}
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temporary sudoers file: %w", err)
	}

	if _, err := runner.Run("sudo", "mv", tmpFile.Name(), sudoersPath); err != nil {
		return fmt.Errorf("failed to move sudoers file: %w", err)
	}
	if _, err := runner.Run("sudo", "chmod", "440", sudoersPath); err != nil {
		return fmt.Errorf("failed to set permissions on sudoers file: %w", err)
	}
	if _, err := runner.Run("sudo", "chown", "root:root", sudoersPath); err != nil {
		return fmt.Errorf("failed to set ownership on sudoers file: %w", err)
	}

	fmt.Println("Sudoers file placed successfully.")
	return nil
}
