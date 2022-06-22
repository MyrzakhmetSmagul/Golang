package src

import (
	"fmt"
	"math"
)

func median(num_Array []int) {
	var medianIndex int
	var medianMean float64
	for i:=0;i<len(num_Array);i++{
		for j:=0;j<len(num_Array)-i-1;j++{
			if num_Array[j] > num_Array[j+1]{
				num_Array[j], num_Array[j+1]=num_Array[j+1],num_Array[j]
			}
		}
	}
	if len(num_Array)%2 != 0 {
		medianIndex = (len(num_Array) + 1) / 2
		fmt.Printf("Median: %d\n", num_Array[medianIndex])
		return
	}
	medianIndex = (len(num_Array)-1) / 2
	medianMean = (float64(num_Array[medianIndex]) + float64(num_Array[medianIndex+1])) / float64(2)
	fmt.Printf("Median %d\n", int(math.Round(medianMean))
}
