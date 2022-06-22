package src

import (
	"fmt"
	"math"
)

func variance(num_Array []int, avg float64) float64 {
	var variance_Num []float64
	var sum float64
	var vrn float64
	for _, elem := range num_Array {
		curr := avg - float64(elem)
		variance_Num = append(variance_Num, curr)
	}
	for _, elem := range variance_Num {
		sum += elem * elem
	}
	vrn = float64(sum) / float64(len(variance_Num))
	fmt.Printf("Variance %d\n", int(math.Round(vrn))
	return vrn
}
