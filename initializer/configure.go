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

// PlaceTorrcConfig places the torrc configuration file at the expected location
func PlaceTorrcConfig(runner runneradapter.CommandRunner) error {
	const torrcPath = "/etc/tor/torrc"

	// Load the template file
	templateContent, err := templates.ReadFile("templates/tor/torrc")
	if err != nil {
		return fmt.Errorf("failed to read torrc template: %w", err)
	}

	// Write to a temporary file
	tmpFile, err := os.CreateTemp("", "torrc-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temporary file for torrc: %w", err)
	}
	defer os.Remove(tmpFile.Name()) // Cleanup in case of failure

	if _, err := tmpFile.Write(templateContent); err != nil {
		return fmt.Errorf("failed to write torrc template to temporary file: %w", err)
	}
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temporary file for torrc: %w", err)
	}

	// Move the temporary file to the final destination
	cmd := []string{"sudo", "mv", tmpFile.Name(), torrcPath}
	if _, err := runner.Run(cmd[0], cmd[1:]...); err != nil {
		return fmt.Errorf("failed to move torrc file: %w", err)
	}

	// Set permissions
	cmd = []string{"sudo", "chmod", "600", torrcPath}
	if _, err := runner.Run(cmd[0], cmd[1:]...); err != nil {
		return fmt.Errorf("failed to set permissions for torrc file: %w", err)
	}

	fmt.Println("Torrc configuration placed successfully.")
	return nil
}

// PlacePrivoxyConfig places the privoxy configuration file at the expected location
func PlacePrivoxyConfig(runner runneradapter.CommandRunner) error {
	const privoxyConfigPath = "/etc/privoxy/config"

	// Load the template file
	templateContent, err := templates.ReadFile("templates/privoxy/config")
	if err != nil {
		return fmt.Errorf("failed to read privoxy config template: %w", err)
	}

	// Write to a temporary file
	tmpFile, err := os.CreateTemp("", "privoxy-config-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temporary file for privoxy config: %w", err)
	}
	defer os.Remove(tmpFile.Name()) // Cleanup in case of failure

	if _, err := tmpFile.Write(templateContent); err != nil {
		return fmt.Errorf("failed to write privoxy config template to temporary file: %w", err)
	}
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temporary file for privoxy config: %w", err)
	}

	// Move the temporary file to the final destination
	cmd := []string{"sudo", "mv", tmpFile.Name(), privoxyConfigPath}
	if _, err := runner.Run(cmd[0], cmd[1:]...); err != nil {
		return fmt.Errorf("failed to move privoxy config file: %w", err)
	}

	// Set permissions
	cmd = []string{"sudo", "chmod", "600", privoxyConfigPath}
	if _, err := runner.Run(cmd[0], cmd[1:]...); err != nil {
		return fmt.Errorf("failed to set permissions for privoxy config file: %w", err)
	}

	fmt.Println("Privoxy configuration placed successfully.")
	return nil
}
