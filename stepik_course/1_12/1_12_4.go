package main

import "fmt"

func main() {
	var workArray [10]uint8

	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}

	//workArray := [10]uint8{99, 151, 137, 71, 117, 187, 20, 93, 187, 67}

	var change1 int
	var change2 int

	for i := 0; i < 3; i++ {
		fmt.Scan(&change1)
		fmt.Scan(&change2)

		workArray[change1], workArray[change2] = workArray[change2], workArray[change1]
	}

	for _, elem := range workArray {
		fmt.Printf("%d ", elem)
	}
}
