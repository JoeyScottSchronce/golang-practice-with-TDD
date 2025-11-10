package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Sum using a slice as an argument", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := Sum(nums)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, nums)
		}
	})

	t.Run("SumAll will return a slice of summed slices", func(t *testing.T) {
		S1 := []int{1, 2, 3, 4, 5}    // 15
		S2 := []int{2, 3, 4, 5, 6, 7} // 27
		S3 := []int{3, 4, 5}          // 12

		got := SumAll(S1, S2, S3)
		want := []int{15, 27, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d when given %v, %v, & %v", got, want, S1, S2, S3)
		}
	})

	t.Run("SumAllTails and return a slice without the heads", func(t *testing.T) {
		S1 := []int{1, 2, 3, 4, 5}    // 15
		S2 := []int{2, 3, 4, 5, 6, 7} // 27
		S3 := []int{3, 4, 5}          // 12

		got := SumAllTails(S1, S2, S3)
		want := []int{14, 25, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d when given %v, %v, & %v", got, want, S1, S2, S3)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

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

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestFind(t *testing.T) {
	t.Run("find the first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		firstEvenNum, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		AssertTrue(t, found)
		AssertEqual(t, firstEvenNum, 2)
	})
}

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		Tam   = Account{Name: "Tam", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(Tam, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(Tam), 175)
}
