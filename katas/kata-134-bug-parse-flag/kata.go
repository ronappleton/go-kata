package kata

import "errors"

var ErrInvalidFlag = errors.New("invalid flag")

// ParseFlag currently treats any non-empty string as true.
func ParseFlag(input string) (bool, error) {
	if input == "" {
		return false, ErrInvalidFlag
	}
	return true, nil
}
