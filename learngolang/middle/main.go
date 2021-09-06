package main

import (
	"fmt"
	"go_course/learngolang/middle/datafile"
	"log"
	"sort"
)

func main() {
	//numbers, err := datafile.GetFloats("data.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var sum float64 = 0
	//
	//for _, number := range numbers {
	//	sum += number
	//}
	//
	//middle := sum / float64(len(numbers))
	//fmt.Println(middle)

	lines, err := datafile.GetStrings("data1.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}

	var names []string
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)

	//for key, value := range counts {
	//	fmt.Println(key, " ", value)
	for _, value := range names {
		fmt.Println(value, " ", counts[value])
	}
}
