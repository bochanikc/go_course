package main

import "fmt"

func main() {
	var max int
	fmt.Scan(&max)
	var array []int
	var a int

	for i := 0; i < max; i++ {
		fmt.Scan(&a)
		array = append(array, a)
	}

	for idx, elem := range array {
		if idx%2 == 0 {
			fmt.Printf("%d ", elem)
		}
	}
}
