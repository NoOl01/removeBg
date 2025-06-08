package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func PipCheck(command string, args ...string) (string, string, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(command, args...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if strings.Contains(stderr.String(), "WARNING: Package(s) not found") {
			return out.String(), stderr.String(), nil
		}
		return out.String(), stderr.String(), err
	}

	return out.String(), stderr.String(), nil
}
