package homework

func sortMapValues(input map[int]string) (result []string) {
	result = make([]string, 0, len(input))
	for i := 0; len(result) < len(input); i++ {
		for k, v := range input {
			if i == k {
				result = append(result, v)
			}
		}
	}
	return result
}
