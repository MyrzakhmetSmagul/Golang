package src

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Open_File(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	if err != nil {
		fmt.Printf("%s\n",err)
		return
	}
	defer file.Close()
	var num_Array []int
	var curr int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curr, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Not valid input. \n%s", err)
			return
		}
		num_Array = append(num_Array, curr)
	}
	mean := average(num_Array)
	median(num_Array)
	vrn := variance(num_Array, mean)
	stDeviation := math.Sqrt(vrn)
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stDeviation))
	return
}
