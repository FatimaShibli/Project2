package builtins

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Pwd(w io.Writer, args ...string) error {
	if len(args) > 0 {
		if args[0] == "-L" || args[0] == "-l" {
			wd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("", err)
			}

			_, errr := fmt.Fprintln(w, wd)

			return errr
		} else {
			wds, err := os.Getwd()
			wd, err := filepath.EvalSymlinks(wds)
			if err != nil {
				return err
			}

			_, errr := fmt.Fprintln(w, wd)
			return errr
		}
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("", err)
		}
		_, errr := fmt.Fprintln(w, wd)
		return errr
	}
}
