package common

import "strings"

func StringStrip(input string) string {
	if input == "" {
		return ""
	}
	return strings.Join(strings.Fields(input), "")
}
