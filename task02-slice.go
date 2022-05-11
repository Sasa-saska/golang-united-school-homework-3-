package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input))
	copy(result, input)
	for i := len(result)/2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}

	return result
}
