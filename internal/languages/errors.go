package languages

import "errors"

var (
	errMissingID   = errors.New("language ID is required")
	errDuplicateID = errors.New("language ID already registered")
)

type errExtensionConflict struct {
	ext    string
	first  string
	second string
}

func (e errExtensionConflict) Error() string {
	return "extension " + e.ext + " already belongs to " + e.first + " (cannot register " + e.second + ")"
}
