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
	if checkTorService() {
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
	if checkPrivoxyService() {
		fmt.Println("- Privoxy Service [OK]")
	} else {
		fmt.Println("- Privoxy Service [MISSING]")
		if fix {
			fmt.Println("  -> Attempting to place Privoxy Service...")

			runner := &runneradapter.RealCommandRunner{}
			PlacePrivoxyServiceFile(runner)
		}
	}

	// Torrc File Check
	if verifyTorrcConfig() {
		fmt.Println("- Torrc config [OK]")
	} else {
		if fix {
			fmt.Println("  -> Attempting to place Tor configuration...")

			runner := &runneradapter.RealCommandRunner{}
			PlaceTorrcConfig(runner)
		}
	}

	// Privoxy Config Check

	// Iptables Configuration Check
	// if isIptablesConfigured() {
	// 	fmt.Println("- Iptables Config [OK]")
	// } else {
	// 	fmt.Println("- Iptables Config [MISSING]")
	// 	if fix {
	// 		fmt.Println("  -> Attempting to configure Iptables...")
	// 		configureIptables()
	// 	}
	// }

	// IPv6 Support Check
	// if isIPv6Enabled() {
	// 	fmt.Println("- IPv6 Support [ENABLED]")
	// } else {
	// 	fmt.Println("- IPv6 Support [DISABLED]")
	// 	if fix {
	// 		fmt.Println("  -> IPv6 support must be manually enabled.")
	// 	}
	// }
}
