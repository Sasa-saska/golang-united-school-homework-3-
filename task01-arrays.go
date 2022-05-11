package homework

func average(input [15]float32) (result float32) {
	var output, counter float32
	for _, v := range input {
		if v != 0 {
			output += v
			counter++
		}
	}
	output = output / counter
	return output
}
