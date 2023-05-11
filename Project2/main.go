package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/FatimaShibli/Project2/Project2/builtins"
	//"github.com/jh125486/CSCE4600/Project2/builtins"
)

func main() {
	exit := make(chan struct{}, 2) // buffer this so there's no deadlock.
	runLoop(os.Stdin, os.Stdout, os.Stderr, exit)

}

func runLoop(r io.Reader, w, errW io.Writer, exit chan struct{}) {
	var (
		input    string
		err      error
		readLoop = bufio.NewReader(r)
	)
	for {
		select {
		case <-exit:
			_, _ = fmt.Fprintln(w, "exiting gracefully...")
			return
		default:
			if err := printPrompt(w); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if input, err = readLoop.ReadString('\n'); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if err = handleInput(w, input, exit); err != nil {
				_, _ = fmt.Fprintln(errW, err)
			}
		}
	}
}

func printPrompt(w io.Writer) error {
	// Get current user.
	// Don't prematurely memoize this because it might change due to `su`?
	u, err := user.Current()
	if err != nil {
		return err
	}
	// Get current working directory.
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// /home/User [Username] $
	_, err = fmt.Fprintf(w, "%v [%v] $ ", wd, u.Username)

	return err
}

func handleInput(w io.Writer, input string, exit chan<- struct{}) error {
	// Remove trailing spaces.
	input = strings.TrimSpace(input)

	// Split the input separate the command name and the command arguments.
	args := strings.Split(input, " ")
	name, args := args[0], args[1:]

	// Check for built-in commands.
	// New builtin commands should be added here. Eventually this should be refactored to its own func.
	switch name {
	case "cd":
		return builtins.ChangeDirectory(args...)
	case "echo":
		return builtins.Echo(args...)
	case "ls":
		return builtins.List(args...)
	case "env":
		return builtins.EnvironmentVariables(w, args...)
	case "exit":
		exit <- struct{}{}
		return nil
	case "sh":
		if len(args) == 0 {
			return errors.New("missing shell command")
		}
		// Call the shCommand function with the remaining arguments
		output := shCommand(args)

		// Print the command output
		fmt.Fprintln(w, output)
		return nil
	case "bash":
		if len(args) == 0 {
			return errors.New("missing script path")
		}
		return executeBashScript(args[0], args[1:])
	}

	return executeCommand(name, args...)
}

func executeCommand(name string, arg ...string) error {
	// Otherwise prep the command
	cmd := exec.Command(name, arg...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}

func executeBashScript(scriptPath string, args []string) error {
	var cmd *exec.Cmd

	// Check if bash is available
	_, err := exec.LookPath("bash")
	if err == nil {
		cmd = exec.Command("bash", scriptPath)
	} else {
		// Fallback to sh if bash is not available
		cmd = exec.Command("sh", scriptPath)
	}

	// Pass arguments to the command
	cmd.Args = append(cmd.Args, args...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
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
