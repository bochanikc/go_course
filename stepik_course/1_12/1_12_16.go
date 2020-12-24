package main

import "fmt"

func main() {
	var max int
	fmt.Scan(&max)
	var array []int
	var a int
	var count int

	for i := 0; i < max; i++ {
		fmt.Scan(&a)
		array = append(array, a)
	}

	for _, elem := range array {
		if elem > 0 {
			count++
		}
	}
	fmt.Print(count)
}
