package utils

import "strings"

func FormatFileName(name string) (string, error) {
	return strings.ReplaceAll(name, " ", ""), nil
}
