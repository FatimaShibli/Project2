package builtins

import (
	"fmt"
	"os"
	"strings"
)

func Echo(args ...string) error {
	var output string
	var end string
	var escape bool
	escape = true
	end = "\n"
	for i, arg := range args {
		if i == 0 && strings.HasPrefix(arg, "-") {
			// check if options are specified
			if strings.Contains(arg, "n") {
				end = ""
			}
			if strings.Contains(arg, "e") {
				if strings.Contains(arg, "E") {
					if strings.Index(arg, "e") < strings.Index(arg, "E") {
						escape = false
					}
				}
			} else if strings.Contains(arg, "E") {
				escape = false
			}
			continue
		}
		if escape {
			arg = interpretEscapes(arg)
		}
		output += arg
	}
	output += end
	if _, err := fmt.Fprint(os.Stdout, output); err != nil {
		return err
	}

	return nil
}
func interpretEscapes(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			i++
			if i >= len(s) {
				break
			}
			switch s[i] {
			case 'a':
				result.WriteRune('\a')
			case 'b':
				result.WriteRune('\b')
			case 'c':
				return result.String()
			case 'e', 'E':
				result.WriteRune('\033')
			case 'f':
				result.WriteRune('\f')
			case 'n':
				result.WriteRune('\n')
			case 'r':
				result.WriteRune('\r')
			case 't':
				result.WriteRune('\t')
			case 'v':
				result.WriteRune('\v')
			case '\\':
				result.WriteRune('\\')
			default:
				result.WriteByte('\\')
				result.WriteByte(s[i])
			}
		} else {
			result.WriteByte(s[i])
		}
	}
	return result.String()
}

