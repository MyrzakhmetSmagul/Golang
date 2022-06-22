package src

import (
	"fmt"
	"math"
)

func average(num_Array []int) float64 {
	var sum int
	var avg float64
	for _, elem := range num_Array {
		sum += elem
	}
	avg = float64(sum) / float64(len(num_Array))
	fmt.Printf("Average: %d\n", int(math.Round(avg)))
	return avg
}
