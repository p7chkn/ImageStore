package utils

import (
	"goImageStore/iternal/validators"
	"strings"
)

func FormatFileName(name string) (string, error) {
	result := strings.ReplaceAll(name, " ", "")
	if err := validators.FileNameValidate(result); err != nil {
		return "", err
	}
	return result, nil
}
