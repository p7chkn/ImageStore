package validators

import (
	"errors"
	"path/filepath"
)

var (
	extensions = [...]string{".png", ".jpg"}
)

func FileNameValidate(name string) error {
	if len(name) > 60 {
		return errors.New("too long file name")
	}
	if !inAllowedExtensions(extensions, filepath.Ext(name)) {
		return errors.New("wrong extension")
	}
	return nil
}

func inAllowedExtensions(s [2]string, target string) bool {
	for _, item := range s {
		if item == target {
			return true
		}
	}
	return false
}
