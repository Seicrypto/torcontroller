package controller

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// ApplyIptablesRulesFactory applies the given iptables rules
func (h *CommandHandler) ApplyIptablesRulesFactory(rules []struct {
	Command string
	Args    []string
}) error {
	for _, rule := range rules {
		h.Logger.Printf("[INFO] Applying rule: %s %s", rule.Command, strings.Join(rule.Args, " "))

		cmd := exec.Command("sudo", append([]string{rule.Command}, rule.Args...)...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			h.Logger.Printf("[ERROR] Failed to apply rule: %v. Error: %s", rule, stderr.String())
			return fmt.Errorf("failed to apply iptables rule: %w", err)
		}
	}
	h.Logger.Println("[INFO] All iptables rules applied successfully.")
	return nil
}

// ClearIptablesRulesFactory removes the given iptables rules
func (h *CommandHandler) ClearIptablesRulesFactory(rules []struct {
	Command string
	Args    []string
}) error {
	for _, rule := range rules {
		h.Logger.Printf("[INFO] Clearing rule: %s %s", rule.Command, strings.Join(rule.Args, " "))

		cmd := exec.Command("sudo", append([]string{rule.Command}, rule.Args...)...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			h.Logger.Printf("[ERROR] Failed to clear rule: %v. Error: %s", rule, stderr.String())
			return fmt.Errorf("failed to clear iptables rule: %w", err)
		}
	}
	h.Logger.Println("[INFO] All iptables rules cleared successfully.")
	return nil
}

// IPv4 rules for applying and clearing
var ipv4RulesApply = []struct {
	Command string
	Args    []string
}{
	{"iptables", []string{"-t", "nat", "-A", "OUTPUT", "-p", "tcp", "--dport", "80", "-j", "REDIRECT", "--to-ports", "8118"}},
	{"iptables", []string{"-t", "nat", "-A", "OUTPUT", "-p", "tcp", "--dport", "443", "-j", "REDIRECT", "--to-ports", "8118"}},
}

var ipv4RulesClear = []struct {
	Command string
	Args    []string
}{
	{"iptables", []string{"-t", "nat", "-D", "OUTPUT", "-p", "tcp", "--dport", "80", "-j", "REDIRECT", "--to-ports", "8118"}},
	{"iptables", []string{"-t", "nat", "-D", "OUTPUT", "-p", "tcp", "--dport", "443", "-j", "REDIRECT", "--to-ports", "8118"}},
}

// IPv6 rules for applying and clearing
var ipv6RulesApply = []struct {
	Command string
	Args    []string
}{
	{"ip6tables", []string{"-t", "nat", "-A", "OUTPUT", "-p", "tcp", "--dport", "80", "-j", "REDIRECT", "--to-ports", "8118"}},
	{"ip6tables", []string{"-t", "nat", "-A", "OUTPUT", "-p", "tcp", "--dport", "443", "-j", "REDIRECT", "--to-ports", "8118"}},
}

var ipv6RulesClear = []struct {
	Command string
	Args    []string
}{
	{"ip6tables", []string{"-t", "nat", "-D", "OUTPUT", "-p", "tcp", "--dport", "80", "-j", "REDIRECT", "--to-ports", "8118"}},
	{"ip6tables", []string{"-t", "nat", "-D", "OUTPUT", "-p", "tcp", "--dport", "443", "-j", "REDIRECT", "--to-ports", "8118"}},
}

// ApplyIptablesIPv4 applies IPv4 rules
func (h *CommandHandler) ApplyIptablesIPv4() error {
	return h.ApplyIptablesRulesFactory(ipv4RulesApply)
}

// ClearIptablesIPv4 clears IPv4 rules
func (h *CommandHandler) ClearIptablesIPv4() error {
	return h.ClearIptablesRulesFactory(ipv4RulesClear)
}

// ApplyIptablesIPv6 applies IPv6 rules
func (h *CommandHandler) ApplyIptablesIPv6() error {
	return h.ApplyIptablesRulesFactory(ipv6RulesApply)
}

// ClearIptablesIPv6 clears IPv6 rules
func (h *CommandHandler) ClearIptablesIPv6() error {
	return h.ClearIptablesRulesFactory(ipv6RulesClear)
}
