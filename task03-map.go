package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	sortedKeys := make([]int, len(input))

	for k := range input {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)

	for k := range sortedKeys {
		result = append(result, input[k])
	}
	return result
}
