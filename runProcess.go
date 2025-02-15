package main

import (
	"fmt"
	"os/exec"
)

func runProcess(name string, args ...string) string {
	result, err := exec.Command(name, args...).Output()

	if err != nil {
		fmt.Println("Uh oh")
		fmt.Println(err)
	}

	return string(result)
}
