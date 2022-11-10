package structx

func assert(length int) {
	if length == 0 {
		panic("values length is 0")
	}
}

// Max
func Max[T Value](values ...T) T {
	assert(len(values))

	var max = values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

// Min
func Min[T Value](values ...T) T {
	assert(len(values))

	var min = values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

// Sum
func Sum[T Value](values ...T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}
