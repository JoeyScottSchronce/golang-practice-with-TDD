package main

func Sum(arr []int) int {
	var result int
	for _, v := range arr {
		result += v
	}
	return result
}

func SumAll(slices ...[]int) []int {
	result := make([]int, len(slices))
	for i, arr := range slices {
		result[i] = Sum(arr)
	}
	return result
}

func SumAllTails(arr ...[]int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		if len(v) == 0 {
			result[i] = 0
		} else {
			result[i] = Sum(v[1:])
		}
	}
	return result
}
