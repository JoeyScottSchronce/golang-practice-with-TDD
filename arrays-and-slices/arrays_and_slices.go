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

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
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

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}
