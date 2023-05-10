package builtins

import (
	"fmt"
	"os"
)

func List(args ...string) error {
	files, err := os.ReadDir(".")
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}
