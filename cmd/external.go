package cmd

import (
	"os"
	"os/exec"
	"strings"
)

func Execute(input string) error {
	// Removing newline character
	input = strings.TrimSuffix(input, "\n")
	args := strings.Fields(input)

	if len(args) != 0 {
		// Checking for built-in commands
		err, flag := executeBuiltIn(args)
		if err != nil {
			return err
		}
		if flag {
			return nil
		}
		// Handlding external commands
		cmd := exec.Command(args[0], args[1:]...)

		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		return cmd.Run()
	}
	return nil
}
