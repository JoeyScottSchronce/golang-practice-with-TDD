package main

func Sum(arr []int) int {
	add := func(y, x int) int { return y + x }
	return Reduce(arr, add, 0)
}

func SumAll(slices ...[]int) []int {

	addAll := func(arr, x []int) []int {
		return append(arr, Sum(x))
	}
	return Reduce(slices, addAll, []int{})
}

func SumAllTails(arr ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(arr, sumTail, []int{})
}

func Reduce[T any](collection []T, f func(T, T) T, initialValue T) T {
	var result = initialValue

	for _, v := range collection {
		result = f(result, v)
	}
	return result
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64

	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}
