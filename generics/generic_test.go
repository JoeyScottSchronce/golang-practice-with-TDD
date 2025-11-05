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

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}

func TestStack(t *testing.T) {
	t.Run("stack of intergers", func(t *testing.T) {

		StackOfInts := NewStack[int]()
		AssertTrue(t, StackOfInts.IsEmpty())

		StackOfInts.Push(32)
		AssertFalse(t, StackOfInts.IsEmpty())

		StackOfInts.Push(64)
		val, _ := StackOfInts.Pop()
		AssertEqual(t, val, 64)

		val2, _ := StackOfInts.Pop()
		AssertEqual(t, val2, 32)

		AssertTrue(t, StackOfInts.IsEmpty())

		StackOfInts.Push(128)
		StackOfInts.Push(255)

		first, _ := StackOfInts.Pop()
		second, _ := StackOfInts.Pop()

		AssertEqual(t, first+second, 383)
	})

	t.Run("Stack of strings", func(t *testing.T) {

		StackOfStrs := NewStack[string]()
		AssertTrue(t, StackOfStrs.IsEmpty())

		StackOfStrs.Push("8")
		AssertFalse(t, StackOfStrs.IsEmpty())

		StackOfStrs.Push("16")
		val, _ := StackOfStrs.Pop()
		AssertEqual(t, val, "16")

		val2, _ := StackOfStrs.Pop()
		AssertEqual(t, val2, "8")

		AssertTrue(t, StackOfStrs.IsEmpty())

		StackOfStrs.Push("World!")
		StackOfStrs.Push("Hello ")

		first, _ := StackOfStrs.Pop()
		second, _ := StackOfStrs.Pop()

		AssertEqual(t, first+second, "Hello World!")
	})
}
