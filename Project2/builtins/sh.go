package builtins

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func Sh(args ...string) error {
	if len(args) == 0 {
		return errors.New("missing shell command")
	}

	output := shCommand(args)

	// Print the command output
	fmt.Println(output)
	return nil
}

func shCommand(args []string) string {
	// Check if args has at least one element
	if len(args) == 0 {
		return "missing shell command"
	}
	// If args has only one element, pass it directly as the argument
	if len(args) == 1 {
		cmd := exec.Command(args[0])
		output, err := cmd.Output()
		if err != nil {
			return err.Error()
		}
		return strings.TrimSpace(string(output))
	}
	// Execute the command passed as parameter
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.Output()

	// Check for errors
	if err != nil {
		return err.Error()
	}

	// Convert the output to a string and return it
	return strings.TrimSpace(string(output))
}
