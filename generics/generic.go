package main

import "testing"

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Push(val T) {
	s.values = append(s.values, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	i := len(s.values) - 1
	v := s.values[i]
	s.values = s.values[:i]
	return v, true
}
