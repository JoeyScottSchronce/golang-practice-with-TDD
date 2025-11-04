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

type StrStack struct {
	values []string
}

type IntStack struct {
	values []int
}

func (s *StrStack) IsEmpty() bool {
	return len(s.values) == 0
}

func (i *IntStack) IsEmpty() bool {
	return len(i.values) == 0
}

func (s *StrStack) Push(val string) {
	s.values = append(s.values, val)
}

func (i *IntStack) Push(val int) {
	i.values = append(i.values, val)
}

func (s *StrStack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	i := len(s.values) - 1
	v := s.values[i]
	s.values = s.values[:i]
	return v, true
}

func (i *IntStack) Pop() (int, bool) {
	if i.IsEmpty() {
		return 0, false
	}

	j := len(i.values) - 1
	v := i.values[j]
	i.values = i.values[:j]
	return v, true
}
