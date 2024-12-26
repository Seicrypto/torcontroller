package iptable

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/Seicrypto/torcontroller/internal/singleton/logger"
)

// applyIptablesRules applies iptables rules for redirecting traffic to port 8118.
func ApplyIptablesRules() error {
	rules := []struct {
		Command string
		Args    []string
	}{
		// IPv4
		{"iptables", []string{"-t", "nat", "-A", "OUTPUT", "-p", "tcp", "--dport", "80", "-j", "REDIRECT", "--to-ports", "8118"}},
		{"iptables", []string{"-t", "nat", "-A", "OUTPUT", "-p", "tcp", "--dport", "443", "-j", "REDIRECT", "--to-ports", "8118"}},
		// IPv6
		// {"ip6tables", []string{"-t", "nat", "-A", "OUTPUT", "-p", "tcp", "--dport", "80", "-j", "REDIRECT", "--to-ports", "8118"}},
		// {"ip6tables", []string{"-t", "nat", "-A", "OUTPUT", "-p", "tcp", "--dport", "443", "-j", "REDIRECT", "--to-ports", "8118"}},
	}

	logger := logger.GetLogger()
	for _, rule := range rules {
		logger.Info(fmt.Sprintf("Applying rule: %s %s", rule.Command, strings.Join(rule.Args, " ")))

		cmd := exec.Command("sudo", append([]string{rule.Command}, rule.Args...)...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			logger.Error(fmt.Sprintf("Failed to apply rule: %v. Error: %s", rule, stderr.String()))
			return fmt.Errorf("failed to apply iptables rule: %w", err)
		}
	}
	logger.Info("All iptables rules applied successfully.")
	return nil
}

// ClearIptablesRules removes iptables rules for redirecting traffic to port 8118.
func ClearIptablesRules() error {
	rules := []struct {
		Command string
		Args    []string
	}{
		// IPv4
		{"iptables", []string{"-t", "nat", "-D", "OUTPUT", "-p", "tcp", "--dport", "80", "-j", "REDIRECT", "--to-ports", "8118"}},
		{"iptables", []string{"-t", "nat", "-D", "OUTPUT", "-p", "tcp", "--dport", "443", "-j", "REDIRECT", "--to-ports", "8118"}},
		// IPv6
		// {"ip6tables", []string{"-t", "nat", "-D", "OUTPUT", "-p", "tcp", "--dport", "80", "-j", "REDIRECT", "--to-ports", "8118"}},
		// {"ip6tables", []string{"-t", "nat", "-D", "OUTPUT", "-p", "tcp", "--dport", "443", "-j", "REDIRECT", "--to-ports", "8118"}},
	}

	logger := logger.GetLogger()
	for _, rule := range rules {
		logger.Info(fmt.Sprintf("Clearing rule: %s %s", rule.Command, strings.Join(rule.Args, " ")))

		cmd := exec.Command("sudo", append([]string{rule.Command}, rule.Args...)...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			logger.Error(fmt.Sprintf("Failed to clear rule: %v. Error: %s", rule, stderr.String()))
			return fmt.Errorf("failed to clear iptables rule: %w", err)
		}
	}
	logger.Info("All iptables rules cleared successfully.")
	return nil
}
