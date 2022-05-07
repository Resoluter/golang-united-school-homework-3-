package homework

func reverse(input []int64) (result []int64) {

	for i := 0; i < len(input); i++ {
		j := len(input) - 1 - i
		result = append(result, input[j])
	}
	return result
}
