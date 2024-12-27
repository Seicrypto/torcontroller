package initializer_test

import (
	"testing"

	"github.com/Seicrypto/torcontroller/initializer"
)

func TestSudoersFileVerify(t *testing.T) {
	mockTemplates := &MockTemplates{
		Files: map[string][]byte{
			"templates/sudoers.d/torcontroller": []byte("mock sudoers content"),
		},
	}

	mockRunner := &MockCommandRunner{
		Output: "Command executed successfully.",
	}

	mockFileSystem := &MockFileSystem{
		Files: map[string]*MockFileInfo{
			"/etc/sudoers.d/torcontroller": {
				mode: 0o440,
				uid:  0,
				gid:  0,
			},
		},
	}

	mockInitializer := initializer.NewInitializer(mockTemplates, mockRunner, mockFileSystem)

	if !mockInitializer.SudoersFileVerify() {
		t.Errorf("expected SudoersFileVerify to return true")
	}
}
