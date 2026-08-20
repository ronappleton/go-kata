package kata

import "errors"

var (
	// ErrInvalidUserID is returned when input ID is blank after trimming spaces.
	ErrInvalidUserID = errors.New("invalid user id")
	// ErrUserNotFound is returned when an ID is not found in the map.
	ErrUserNotFound = errors.New("user not found")
)

// Find User Name — TODO: implement as specified in README.md
//
// Signature:
//
//	func FindUserName(users map[string]string, id string) (string, error)
func FindUserName(users map[string]string, id string) (string, error) {
	// TODO: your implementation here
	return "", nil
}
