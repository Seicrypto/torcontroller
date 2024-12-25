package runneradapter_test

import "errors"

type MockCommandRunner struct {
	ExpectedCommands map[string]string
}

func (m *MockCommandRunner) Run(name string, args ...string) (string, error) {
	command := name + " " + combineArgs(args)
	if output, ok := m.ExpectedCommands[command]; ok {
		return output, nil
	}
	return "", errors.New("unexpected command: " + command)
}

func combineArgs(args []string) string {
	result := ""
	for _, arg := range args {
		result += arg + " "
	}
	return result[:len(result)-1]
}
