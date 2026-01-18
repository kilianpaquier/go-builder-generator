/*
Package testutils provides simple utility functions for testing.
*/
package testutils

import (
	"errors"
	"strings"
	"testing"
)

// require wraps testing.TB to provide FailNow on Errorf.
type require struct {
	testing.TB
}

// Error logs the args and fails the test immediately.
func (r require) Error(args ...any) {
	r.Helper()
	r.Log(args...)
	r.FailNow()
}

// Errorf logs the formatted error message and fails the test immediately.
func (r require) Errorf(format string, args ...any) {
	r.Helper()
	r.Logf(format, args...)
	r.FailNow()
}

// Require returns a testing.TB that fails immediately on errors.
func Require(t testing.TB) testing.TB {
	t.Helper()
	return require{TB: t}
}

// Contains asserts that str contains substr.
func Contains(t testing.TB, str, substr string) {
	t.Helper()
	if !strings.Contains(str, substr) {
		t.Errorf("expected '%s' to contain '%s'", str, substr)
	}
}

// Equal asserts that expected and actual are equal.
func Equal[T comparable](t testing.TB, expected, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

// Error asserts that err is not nil.
func Error(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Error("expected error, got nil")
	}
}

// ErrorIs asserts that err is or wraps target.
func ErrorIs(t testing.TB, err, target error) {
	t.Helper()
	if !errors.Is(err, target) {
		t.Errorf("expected error '%s', got '%s'", target, err)
	}
}

// NoError asserts that err is nil.
func NoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected no error, got '%s'", err)
	}
}
