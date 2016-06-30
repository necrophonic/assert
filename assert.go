package assert

import "testing"

// ErrorExists asserts whether a given error exits and throws a test error if it doesn't.
func ErrorExists(err error, t *testing.T) {
	if err == nil {
		t.Error("expected an error but got none")
	}
}

// ErrorEquals asserts whether the given error is non-nil and matches the given string
func ErrorEquals(err error, expected string, t *testing.T) {
	if err == nil {
		t.Errorf("expected error to be '%s' but got none", expected)
	} else {
		if err.Error() != expected {
			t.Errorf("expected error to be '%s' but got '%s'", expected, err.Error())
		}
	}
}
