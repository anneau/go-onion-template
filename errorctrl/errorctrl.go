// Package errctrl is intended to simplify error checking and assignments to temporary variables.
package errctrl

// Must panics if the second arg is an error.
//   - Allowed in production code only if it never fails.
//   - Don't use in production code if it is likely to fail. Use only in tests.
func Must[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}
	return result
}

// MustExec panics if the first arg is an error.
//   - Allowed in production code only if it never fails.
//   - Don't use in production code if it is likely to fail. Use only in tests.
func MustExec(err error) {
	if err != nil {
		panic(err)
	}
}

// Ignore indicates ignoring errors.
// This is intended to be combined with `defer` calls.
func Ignore(fn func() error) {
	_ = fn()
}
