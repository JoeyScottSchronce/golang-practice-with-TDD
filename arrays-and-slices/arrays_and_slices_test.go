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
