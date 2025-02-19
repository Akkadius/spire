package filepathcheck

import (
	"errors"
	"strings"
)

func IsValid(input string) error {
	if strings.Contains(input, "..") {
		return errors.New("path traversal detected")
	}

	return nil
}
