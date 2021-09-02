package main

import (
	"fmt"
	"go_course/learngolang/middle/datafile"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	middle := sum / float64(len(numbers))
	fmt.Println(middle)
}
