package util

import (
	"fmt"
)

// Color wraps input around ANSI color escape codes
func Color(input interface{}, color int) string {
	if color == 0 {
		return fmt.Sprintf("%s", input)
	}

	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", color, input)
}
