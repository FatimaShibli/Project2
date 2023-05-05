package builtins

import (
	"fmt"
	"os"
)

func echo(args ...string) error {
	var output string
	for i, arg := range args {
		if i > 0 {
			output += " "
		}
		output += arg
	}

	if _, err := fmt.Fprintln(os.Stdout, output); err != nil {
		return err
	}

	return nil
}
