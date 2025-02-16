package main

import (
	"os/exec"
)

// No longer used, originally I retrieved metrics by running CLI commands
// but switched to using gopsutil.
// Keeping for potential future use
func runProcess(name string, args ...string) string {
	result, err := exec.Command(name, args...).Output()

	if err != nil {
		panic(err)
	}

	return string(result)
}
