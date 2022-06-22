package main

import (
	"fmt"
	"main/src"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter only the file name \"txt\" after the program name!")
		return
	}
	src.Open_File(os.Args[1])
}
