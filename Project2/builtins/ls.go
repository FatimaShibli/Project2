package builtins

import (
	"fmt"
	"os"
	"path/filepath"
)

// List returns a list of files in the current directory.
func List(args ...string) error {
	files, err := os.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read directory: %v", err)
	}
	for _, file := range files {
		fmt.Println(filepath.Join(".", file.Name()))
	}
	return nil
}
