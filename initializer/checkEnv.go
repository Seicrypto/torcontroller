package initializer

import (
	"fmt"

	runneradapter "github.com/Seicrypto/torcontroller/internal/services/runnerAdapter"
)

func CheckEnvironment(fix bool) {
	fmt.Println("Environment Check Report:")

	// Sudoer File Check
	if sudoersFileVerify() {
		fmt.Println("- Sudoers File [OK]")
	} else {
		fmt.Println("- Sudoers File [MISSING]")
		if fix {
			fmt.Println("  -> Attempting to place Sudoers File...")

			runner := &runneradapter.RealCommandRunner{}
			PlaceSudoersFile(runner)
		}
	}

	// Tor Service File Check
	if verifyTorService() {
		fmt.Println("- Tor Service [OK]")
	} else {
		fmt.Println("- Tor Service [MISSING]")
		if fix {
			fmt.Println("  -> Attempting to place Tor Service...")

			runner := &runneradapter.RealCommandRunner{}
			PlaceTorServiceFile(runner)
		}
	}

	// Privoxy Service File Check
	if verifyPrivoxyService() {
		fmt.Println("- Privoxy Service [OK]")
	} else {
		fmt.Println("- Privoxy Service [MISSING]")
		if fix {
			fmt.Println("  -> Attempting to place Privoxy Service...")

			runner := &runneradapter.RealCommandRunner{}
			PlacePrivoxyServiceFile(runner)
		}
	}

}
