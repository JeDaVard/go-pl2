package main

func countAvg(values []int64) int {
	if len(values) == 0 {
		return 0
	}
	var sum int64
	for _, value := range values {
		sum += value
	}
	return int(sum) / len(values)
}
