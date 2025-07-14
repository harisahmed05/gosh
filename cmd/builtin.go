package cmd

import (
	"fmt"
	"os"
)

func executeBuiltIn(args []string) (error, bool) { // bool indicates if any built-in command was entered by user
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			homedir, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("failed to get home directory: %w", err), true
			}
			return os.Chdir(homedir), true
		}
		return os.Chdir(args[1]), true
	case "exit":
		os.Exit(0)
	}
	return nil, false
}
