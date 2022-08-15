//go:build go1.12
// +build go1.12

package shareddefaults

import (
	"os"
	"os/user"
)

func userHomeDir() string {
	home, _ := os.UserHomeDir()

	if len(home) > 0 {
		return home
	}

	currUser, _ := user.Current()
	if currUser != nil {
		home = currUser.HomeDir
	}

	return home
}
