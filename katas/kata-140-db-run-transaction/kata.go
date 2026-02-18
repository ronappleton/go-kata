package kata

import "errors"

var ErrRollback = errors.New("rollback")

// RunInTransaction — TODO: implement as specified in README.md
//
// Signature:
//
//	func RunInTransaction(begin func() error, commit func() error, rollback func() error, run func() error) error
func RunInTransaction(begin func() error, commit func() error, rollback func() error, run func() error) error {
	// TODO: your implementation here
	return nil
}
