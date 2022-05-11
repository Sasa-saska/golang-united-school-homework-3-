package homework

import "sort"

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input))
	copy(result, input)
	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})

	return result
}
