package utils

func ConvertToIntSlice(answer interface{}) ([]int, bool) {
	interfaceSlice, ok := answer.([]interface{})
	if !ok {
		return nil, false
	}

	intSlice := make([]int, len(interfaceSlice))
	for i, v := range interfaceSlice {
		if floatVal, ok := v.(float64); ok {
			intSlice[i] = int(floatVal)
		} else {
			return nil, false
		}
	}
	return intSlice, true
}

func CalculateCorrectAdjacentPairs(correctOrder, submittedOrder []int) float64 {
	if len(correctOrder) != len(submittedOrder) {
		return 0
	}

	correctPositions := make(map[int]int)
	for i, val := range correctOrder {
		correctPositions[val] = i
	}

	correctPairs := 0
	for i := 0; i < len(submittedOrder)-1; i++ {
		if correctPositions[submittedOrder[i]] < correctPositions[submittedOrder[i+1]] {
			correctPairs++
		}
	}

	return float64(correctPairs)
}
