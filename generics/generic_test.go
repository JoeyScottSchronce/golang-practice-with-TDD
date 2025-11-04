package main

import (
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on intergers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("assert on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "not hello")
	})
}

// next test here
